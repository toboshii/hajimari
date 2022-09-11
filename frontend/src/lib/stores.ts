import { writable, derived } from 'svelte/store';

export const appQuery = writable('');
export const apps = writable([]);
export const filteredApps = derived(
    [appQuery, apps], 
    ([$appQuery, $apps]) => $apps.reduce((r, o) => {
        var apps = o.apps.filter(a => a.name.toLowerCase().includes($appQuery))
        if (apps.length) r.push(Object.assign({}, o, apps.length && {apps}))
        return r;
    }, [])
);