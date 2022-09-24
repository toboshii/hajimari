import { writable, derived } from 'svelte/store';

export const appQuery = writable('');
export const apps = writable([]);
export const filteredApps = derived(
    [appQuery, apps],
    ([$appQuery, $apps]) => $apps.reduce((r, o) => {
        var apps = o.apps.filter(a => a.name.toLowerCase().includes($appQuery))
        if (apps.length) r.push(Object.assign({}, o, apps.length && { apps }))
        return r;
    }, [])
);
export const themes = writable([
    {
        'name': 'blackboard',
        'backgroundColor': '#1a1a1a',
        'primaryColor': '#FFFDEA',
        'accentColor': '#5c5c5c'
    },
    {
        'name': 'gazette',
        'backgroundColor': '#F2F7FF',
        'primaryColor': '#000000',
        'accentColor': '#5c5c5c'
    },
    {
        'name': 'espresso',
        'backgroundColor': '#21211F',
        'primaryColor': '#D1B59A',
        'accentColor': '#4E4E4E'
    },
    {
        'name': 'cab',
        'backgroundColor': '#F6D305',
        'primaryColor': '#1F1F1F',
        'accentColor': '#424242'
    },
    {
        'name': 'cloud',
        'backgroundColor': '#f1f2f0',
        'primaryColor': '#35342f',
        'accentColor': '#37bbe4'
    },
    {
        'name': 'lime',
        'backgroundColor': '#263238',
        'primaryColor': '#AABBC3',
        'accentColor': '#aeea00'
    },
    {
        'name': 'tron',
        'backgroundColor': '#242B33',
        'primaryColor': '#EFFBFF',
        'accentColor': '#6EE2FF'
    },
    {
        'name': 'blues',
        'backgroundColor': '#2B2C56',
        'primaryColor': '#EFF1FC',
        'accentColor': '#6677EB'
    },
    {
        'name': 'passion',
        'backgroundColor': '#f5f5f5',
        'primaryColor': '#12005e',
        'accentColor': '#8e24aa'
    },
    {
        'name': 'chalk',
        'backgroundColor': '#263238',
        'primaryColor': '#AABBC3',
        'accentColor': '#FF869A'
    },
    {
        'name': 'paper',
        'backgroundColor': '#F8F6F1',
        'primaryColor': '#4C432E',
        'accentColor': '#AA9A73'
    },
    {
        'name': 'horizon',
        'backgroundColor': '#232530',
        'primaryColor': '#FAB795',
        'accentColor': '#E95678'
    }
]);