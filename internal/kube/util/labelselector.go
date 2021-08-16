package util

import (
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func labelSelectorAsMap(ps *metav1.LabelSelector) (map[string]string, error) {
	if ps == nil {
		return nil, nil
	}

	selector := map[string]string{}
	for k, v := range ps.MatchLabels {
		selector[k] = v
	}

	for _, expr := range ps.MatchExpressions {
		switch expr.Operator {
		case metav1.LabelSelectorOpIn:
			if len(expr.Values) != 1 {
				return selector, errors.New("operator " + string(expr.Operator) + " without a single value cannot be converted into the old label selector format")
			}
			// Should we do anything in case this will override a previous key-value pair?
			selector[expr.Key] = expr.Values[0]
		case metav1.LabelSelectorOpNotIn, metav1.LabelSelectorOpExists, metav1.LabelSelectorOpDoesNotExist:
			return selector, errors.New("operator " + string(expr.Operator) + " cannot be converted into the old label selector format")
		default:
			return selector, errors.New(string(expr.Operator) + " is not a valid selector operator")
		}
	}
	return selector, nil
}
