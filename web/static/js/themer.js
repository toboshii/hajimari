const setValue = (property, value) => {
    if (value) {
        document.documentElement.style.setProperty(`--${property}`, value);

        const input = document.querySelector(`#${property}`);
        if (input) {
            value = value.replace('px', '');
            input.value = value;
        }
    }
};

const setValueFromLocalStorage = property => {
    let value = localStorage.getItem(property);
    setValue(property, value);
};

const setTheme = options => {
    for (let option of Object.keys(options)) {
        const property = option;
        const value = options[option];

        setValue(property, value);
        localStorage.setItem(property, value);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    setValueFromLocalStorage('color-background');
    setValueFromLocalStorage('color-text-pri');
    setValueFromLocalStorage('color-text-acc');
    setValueFromLocalStorage('color-status-online');
    setValueFromLocalStorage('color-status-offline');
});

const dataThemeButtons = document.querySelectorAll('[data-theme]');

for (let i = 0; i < dataThemeButtons.length; i++) {
    dataThemeButtons[i].addEventListener('click', () => {
        const theme = dataThemeButtons[i].dataset.theme;

        switch (theme) {
            case 'blackboard':
                setTheme({
                    'color-background': '#1a1a1a',
                    'color-text-pri': '#FFFDEA',
                    'color-text-acc': '#5c5c5c',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'gazette':
                setTheme({
                    'color-background': '#F2F7FF',
                    'color-text-pri': '#000000',
                    'color-text-acc': '#5c5c5c',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'espresso':
                setTheme({
                    'color-background': '#21211F',
                    'color-text-pri': '#D1B59A',
                    'color-text-acc': '#4E4E4E',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'cab':
                setTheme({
                    'color-background': '#F6D305',
                    'color-text-pri': '#1F1F1F',
                    'color-text-acc': '#424242',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'cloud':
                setTheme({
                    'color-background': '#f1f2f0',
                    'color-text-pri': '#35342f',
                    'color-text-acc': '#37bbe4',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'lime':
                setTheme({
                    'color-background': '#263238',
                    'color-text-pri': '#AABBC3',
                    'color-text-acc': '#aeea00',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'white':
                setTheme({
                    'color-background': '#ffffff',
                    'color-text-pri': '#222222',
                    'color-text-acc': '#dddddd',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'tron':
                setTheme({
                    'color-background': '#242B33',
                    'color-text-pri': '#EFFBFF',
                    'color-text-acc': '#6EE2FF',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'blues':
                setTheme({
                    'color-background': '#2B2C56',
                    'color-text-pri': '#EFF1FC',
                    'color-text-acc': '#6677EB',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'passion':
                setTheme({
                    'color-background': '#f5f5f5',
                    'color-text-pri': '#12005e',
                    'color-text-acc': '#8e24aa',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'chalk':
                setTheme({
                    'color-background': '#263238',
                    'color-text-pri': '#AABBC3',
                    'color-text-acc': '#FF869A',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

            case 'paper':
                setTheme({
                    'color-background': '#F8F6F1',
                    'color-text-pri': '#4C432E',
                    'color-text-acc': '#AA9A73',
                    'color-status-online': '#00FF00',
                    'color-status-offline': '#FF0000'
                });
                return;

        }
    })
}
