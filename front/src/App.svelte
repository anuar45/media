<script>
	import { onMount } from 'svelte';
	import { getMediaItems } from "./api/media";
  import Image from './Image.svelte'
  import Directory from './Directory.svelte'
  import Video from './Video.svelte'

  let mediaItems = [];
  let currentPath = "/";
  $ : parentPath = currentPath === "/" ? currentPath : currentPath.replace(/\/.*$/, '');

	onMount(async () => {
    const res = await getMediaItems(currentPath);
    mediaItems = res;
    console.log(mediaItems);
  });

	export const changeDir = async (dirPath) => {
	  const res = await getMediaItems(dirPath);
    mediaItems = res;
    console.log(dirPath);
    console.log(mediaItems);
	};

</script>


<main>
	<div id="gallery" class="container">
    <Directory dirName=".." dirPath={parentPath} changeDir={changeDir}/>
		{#each mediaItems as item}
      {#if item.type == "directory"}
        <Directory dirName={item.name} dirPath={item.path} changeDir={changeDir}/>
      {/if}
		{/each}
		{#each mediaItems as item}
      {#if item.type == "image"}
        <Image imageName={item.name} imageUrl={item.path}/>
      {/if}
		{/each}
		{#each mediaItems as item}
      {#if item.type == "video"}
        <Video videoName={item.name} videoUrl={item.path}/>
      {/if}
		{/each}
	</div>
</main>


<style>
main {
  font-family: Verdana, sans-serif;
  margin: 0;
  background-color:#282c34;
}

#gallery {
  display: flex;
  /* line-height: 0; */
  /* column-count: 5; */
  /* column-gap: 5px; */
  flex-wrap: wrap;
}

</style>
