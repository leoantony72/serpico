<template>
  <div class="encryption">
    <h1>Find Duplicates</h1>
  </div>
  <div class="upload-section">
    <h3>Choose Folder to Scan</h3>
    <div class="button-container">
      <button class="choose-button" @click="selectFile">
        {{ buttonText }}
      </button>
      <button class="upload-button organize-button" @click="encrypt">
        Scan
      </button>
    </div>
    <input
      type="file"
      ref="fileInput"
      @change="handleFileChange"
      style="display: none"
    />
  </div>
</template>

<script setup>
import { ref } from "vue";
const fileInput = ref(null);
const files = ref();
const buttonText = ref("Choose File");

const encrypt = () => {

}

const selectFile = () => {
  fileInput.value.click();
};
const handleFileChange = (event) => {
  const file = event.target.files[0];
  // Handle the selected file
  if (file) {
    files.value = file;
    buttonText.value = file.name; // Update button text to file name
  }
};
</script>

<style>
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
</style>
