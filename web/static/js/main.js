document.addEventListener('DOMContentLoaded', () => {
    var editor = CodeMirror.fromTextArea(document.getElementById("editor"), {
        mode: "yaml",
        theme: "dracula",
        lineNumbers: true,
        tabsize: 2,
        indentWithTabs: false,
        foldGutter: true,
        gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter", "CodeMirror-lint-markers"],
        lint: {
            "getAnnotations": self.yaml_validator
        }
    });

    editor.setOption("extraKeys", {
        Tab: function(cm) {
            var spaces = Array(cm.getOption("indentUnit") + 1).join(" ");
            cm.replaceSelection(spaces);
        }
    });

    fetch(location.href.replace(location.hash, "").replace(/\/$/, "") + '/config')
        .then(response => response.json())
        .then(data => editor.setValue(jsyaml.dump(data)))

    var saveButton = document.getElementById('save-config');

    saveButton.addEventListener('click', function() {
        console.log(JSON.stringify(jsyaml.load(editor.getValue())));
        var method = location.pathname == "/" ? 'post' : 'put';
        fetch(location.href.replace(location.hash,""), {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsyaml.load(editor.getValue()))
        }).then(response => response.json())
        .then(function(data) {
            if (data.error) {

            } else {
                location.assign(location.origin + "/" + data.id);
            }
        });
    }, false);
});