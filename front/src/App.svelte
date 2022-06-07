<script>
	import { onMount } from 'svelte';
	import { getMediaItems } from "./api/media";
  import Image from './Image.svelte'
  import Directory from './Directory.svelte'
  import Video from './Video.svelte'
  import Unknown from './Unknown.svelte'

  let mediaItems = [];
  let currentPath = "/";
  let currentBase = "";
  let parentPath = "";
  $ : parentPath = currentPath === "" ? currentPath : currentPath.replace(/\/\w*$/, '');
  $ : currentBase = currentPath === "" ? "" : currentBase;

	onMount(async () => {
    const res = await getMediaItems(currentBase, currentPath);
    mediaItems = res;
    console.log(mediaItems);
  });

	export const changeDir = async (base, dirPath) => {
    currentBase = base
    currentPath = dirPath
	  const res = await getMediaItems(base, dirPath);
    mediaItems = res;
    console.log("parentPath", parentPath)
    console.log("currentPath", currentPath)
    console.log("dirPath", dirPath);
    console.log(mediaItems);
	};

</script>


<main>
	<div id="gallery" class="container">
    {#if currentBase !== ""}
    <Directory dirName=".." dirBase={currentBase} dirPath={parentPath} changeDir={changeDir}/>
		{/if}
    {#each mediaItems as item}
      {#if item.type == "directory"}
        <Directory dirName={item.name} dirBase={item.base} dirPath={item.path} changeDir={changeDir}/>
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
    {#each mediaItems as item}
      {#if item.type == "unknown"}
        <Unknown itemName={item.name}/>
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
