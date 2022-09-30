import { error } from '@sveltejs/kit';
import { api } from '$lib/api';
import type { PageLoad } from './$types';

type Startpage = {
    id: string;
    name: string;
    title: string;
    theme: string;
    lightTheme: string;
    darkTheme: string;
    customThemes: Array<Record<string, string>>;
    showGreeting: boolean;
    showAppGroups: boolean;
    showAppUrls: boolean;
    showAppInfo: boolean;
    showAppStatus: boolean;
    showBookmarkGroups: boolean;
    showGlobalBookmarks: boolean;
    searchProviders: Array<Record<string, string>>;
    bookmarks: any;
};

export const load: PageLoad = async ({ fetch, params }) => {
    const { slug } = params;

    const [startpage, apps, bookmarks] = await Promise.all([
        api(fetch, 'GET', `startpage/${slug}`),
        api(fetch, 'GET', 'apps'),
        api(fetch, 'GET', 'bookmarks')
    ]);

    if (await startpage.status !== 200) {
        let data = await startpage.json();
        throw error(startpage.status, data.status);
    }

    return {
        startpage: (await startpage.json() as Startpage),
        apps: (await apps.json()),
        globalBookmarks: (await bookmarks.json()),
        slug
    }

};