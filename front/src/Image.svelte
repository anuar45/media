<script>
	export let imageUrl;
	export let imageName;
  // export let imageNum;
</script>

<style>

img {
   width: auto !important;
   height: 100px !important;
}

.image {
  height: 100px;
  padding: 5px
  /*margin-bottom:5px;*/
}

img.hover-shadow {
  transition: 0.3s;
}

.hover-shadow:hover {
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.cursor {
  cursor: pointer;
}
</style>

<div class="image">
  <img src={imageUrl} alt={imageName} class="hover-shadow cursor">
</div>

 
