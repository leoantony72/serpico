<template>
  <div class="encryption">
    <h1>Find Duplicates</h1>
  </div>
  <div class="upload-section">
    <h3>Choose Folder to Scan</h3>
    <div class="button-container">
      <button class="choose-button" @click="SelectFolder()">
        Choose folder
      </button>
      <button class="upload-button organize-button" @click="findDuplicates">
        Scan
      </button>
      <button
        v-if="flatDuplicates.length > 0"
        class="upload-button organize-button"
        @click="deleteDuplicates"
      >
        Delete
      </button>
    </div>
  </div>
  <div class="container" v-if="flatDuplicates.length > 0">
    <div class="image-grid">
      <div
        v-for="(image, index) in flatDuplicates"
        :key="index"
        class="image-item"
      >
        <img :src="getFileUrl(image)" v-if="isImageFile(image)" />
        <img :src="fileIcon()" v-else />
        <p class="file-name">{{ getFileName(image) }}</p>
      </div>
    </div>
  </div>
  <div class="message" v-if="showNoDuplicatesMessage">
    <p>No duplicate files found.</p>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { FindDuplicates } from "/wailsjs/go/main/App";
import { DeleteFiles } from "/wailsjs/go/main/App";
import { SelectDirectoryDuplicate } from "/wailsjs/go/main/App";

const folderPaths = ref([]);
const duplicates = ref(null);
const flatDuplicates = ref([]);
const selectedFolder = ref("");
const showNoDuplicatesMessage = ref(false);
const imageExtensions = ["jpg", "jpeg", "png"];

const SelectFolder = async () => {
  try {
    const path = await SelectDirectoryDuplicate();
    if (path) {
      // const name = path.split(/[\\/]/).pop();
      folderPaths.value.push(path);
      selectedFolder.value = path;
    }
  } catch (error) {
    console.error("Error selecting folder:", error);
  }
};
const isImageFile = (file) => {
  const extension = file.split(".").pop().toLowerCase();
  return imageExtensions.includes(extension);
};

const getFileUrl = (file) => {
  const relativePath = file
    .replace(selectedFolder.value, "")
    .replace(/\\/g, "/");
  return `http://localhost:8080/assets${relativePath}`;
};

const fileIcon = () => {
  return `/src/assets/images/default.png`;
};

const getFileName = (file) => {
  return file.split(/[/\\]/).pop();
};

const deleteDuplicates = async () => {
  try {
    await DeleteFiles(flatDuplicates.value);
    flatDuplicates.value = [];
    duplicates.value = [];
  } catch (error) {
    console.error("Error deleting files:", error);
  }
};

const findDuplicates = async () => {
  try {
    const result = await FindDuplicates(selectedFolder.value);
    duplicates.value = result;
    flatDuplicates.value = Object.values(result).flat();
    // if Object.keys(result)
    if (flatDuplicates.value.length === 0) {
      showNoDuplicatesMessage.value = true;
    } // Reset message flag
    console.log(result);
    console.log(flatDuplicates);
  } catch (error) {
    console.error("Error finding duplicates:", error);
  }
};

const fileTypeImages = {
  default: "public/default-images/default.png",
  jpg: "public/default-images/images.png",
  jpeg: "public/default-images/images.png",
  png: "public/default-images/images.png",
  txt: "public/default-images/text.jpg",
  pdf: "public/default-images/pdf.jpg",
  mp3: "public/default-images/audio.png",
  mp4: "public/default-images/mp4.jpg",
  // Add more file type mappings here
};

const handleImageError = (event) => {
  console.error("Image loading error:", event.target.src);
};
</script>

<style>
.message p {
  margin-left: 160px;
  color: #1b2821;
  font-family: "Outfit";
}
.encryption {
  height: 100px;
  display: flex;
  justify-content: center;
  margin-top: 50px;
  text-align: center;
}

.encryption h1 {
  color: #1b2821;
  font-family: "Outfit";
  font-size: 45px;
}
.button-container {
  display: flex;
  justify-content: flex-start; /* Align buttons to the left */
  gap: 10px;
}

.upload-section h3 {
  color: #1b2821;
  font-family: "Outfit";
  font-size: 25px;
}

.choose-button {
  color: #1b2821;
  font-family: "Outfit";
  font-weight: 500;
  font-size: 15px;
  background-color: #cfcfcf;
  opacity: 80%;
  margin-top: 20px;
  width: 325px;
  height: 35px;
  border-radius: 8px;
  border: none;
  display: flex; /* Add this line */
  align-items: center; /* Add this line */
  padding-left: 10px; /* Add this line to adjust padding */
}

.choose-button button {
  justify-content: left;
  text-align: left;
  width: 100%; /* Ensure the button takes full width */
  background: none; /* Remove default button background */
  border: none; /* Remove default button border */
  padding: 0; /* Remove default button padding */
}

.upload-button {
  color: #fafefc;
  font-family: "Outfit";
  font-weight: 600;
  font-size: 15px;
  background-color: #4c956c;
  /* opacity: 25%; */
  margin-top: 20px;
  width: 155px;
  height: 35px;
  border-radius: 8px;
  border: none;
}

.upload-button:hover {
  background-color: #3b7955;
}

.container {
  margin-top: 50px;
  max-width: 800px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(4, 0.2fr);
  gap: 5px;
}

.image-group {
  display: flex;
  justify-content: center;
  align-items: center;
}

.image-item img {
  width: 100px;
  height: 100px;
}

.image-item p {
  max-width: 80px;
  max-height: 30px;
  overflow: hidden;
  margin-left: 10px;
  font-family: "Outfit";
  font-size: 13px;
}
.small {
  width: 100px; /* Set the desired width */
  height: 100px; /* Set the desired height */
  object-fit: cover;
  border-radius: 5px;
}

.small:hover {
  transform: scale(1.03, 1.03);
}
</style>
