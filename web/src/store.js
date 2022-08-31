import { writable, derived } from 'svelte/store';

/** Store for your data. 
This assumes the data you're pulling back will be an array.
If it's going to be an object, default this to an empty object.
**/
export const apiData = writable({});

/** Data transformation.
**/
export const name = derived(apiData, ($apiData) => {
  if ($apiData.name){
    return $apiData.name;
  }
  return "You";
});

export const apps = writable([]);

export const bookmarks = writable([]);