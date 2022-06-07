<script>
	export let dirName;
	export let dirPath;
  export let dirBase;
	export let changeDir;
</script>

<style>

img {
  width: auto;
  height: 75%;
}

.directory {
  height: 100px;
  padding: 5px;
}

.caption {
  text-align: center;
  color: white;
}

p {
  margin-top: 0px;
  margin-bottom: 0px;
}

.cursor {
  cursor: pointer;
}
</style>

<div class="directory">
  <img class="cursor" src="./public/directory.png" alt={dirName} on:click={changeDir(dirBase, dirPath)}>
  <div class="caption">
    <p>{dirName}</p>
  </div>
</div>

