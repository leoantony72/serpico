<script setup>
import { ref } from "vue";

import { OrganizeFolder } from "/wailsjs/go/main/App";
import { SelectDirectory } from "/wailsjs/go/main/App";

const folderPaths = ref([]);

const organizeF = async () => {
  try {
    for (let i = 0; i < folderPaths.value.length; i++) {
      const folder = folderPaths.value[i];
      await OrganizeFolder(folder.path, folder.organizeBy);
      console.log(folder.path, folder.organizeBy);
    }
    while (folderPaths.value.length > 0) {
      deleteFolder(0);
    }
  } catch (err) {
    console.log(err);
  }
};

const selectFolder = async () => {
  try {
    const path = await SelectDirectory();
    if (path) {
      const name = path.split(/[\\/]/).pop();
      folderPaths.value.push({ name: name, path, organizeBy: "File Type" });
    }
  } catch (error) {
    console.error("Error selecting folder:", error);
  }
};

const organizeOptions = ["Year", "Month", "File Type"];

const deleteFolder = (index) => {
  folderPaths.value.splice(index, 1);
};
</script>

<template>
  <div class="organizer">
    <h1>Organize Your files</h1>
  </div>
  <div class="upload-section">
    <h3>Upload Your Folder</h3>
    <div class="button-container">
      <button class="upload-button" @click="selectFolder">Choose Folder</button>
      <button class="upload-button organize-button" @click="organizeF">
        Organize
      </button>
    </div>
  </div>
  <div class="folder-table" style="overflow-x: auto">
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Path</th>
          <th>Organize by</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in folderPaths" :key="index">
          <td class="ellipsis" data-label="Name">{{ item.name }}</td>
          <td class="ellipsis" data-label="Path">{{ item.path }}</td>
          <td class="organize-by" data-label="OrganizeBy">
            <select
              v-model="item.organizeBy"
              @change="updateOrganizeBy(index, item.organizeBy)"
            >
              <option
                v-for="option in organizeOptions"
                :key="option"
                :value="option"
              >
                {{ option }}
              </option>
            </select>
          </td>
          <td data-label="Action">
            <button class="remove-folder" @click="deleteFolder(index)">
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style>
.organize-by select {
  outline: none;
  border: none;
  font-family: "Outfit";
}
table {
  width: 100%;
  border: none;
}

th,
td {
  text-align: left;
  padding: 8px;
  /* max-width: 100px; */
}
.folder-table {
  box-sizing: border-box;
  margin-left: 160px;
  margin-right: 10px;
  /* width: calc(100% - 150px); */
}

th {
  color: #1a2821;
  opacity: 62%;
  font-family: "Outfit";
}
td {
  color: #1b2821;
  font-family: "Outfit";
}

.ellipsis {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.organize-button {
  /* display: flex; */
  justify-content: flex-end;
  align-items: flex-end;
  /* margin-left: 450px; */
}
.button-container {
  display: flex;
  justify-content: flex-start; /* Align buttons to the left */
  gap: 10px;
}
.remove-folder {
  border: none;
  background-color: rgb(241, 91, 91);
  border-radius: 6px;
  width: 80px;
  height: 30px;
  font-family: "Outfit";
  color: #fafefc;
  /* transition:  0.1s ease;  */
}
.remove-folder:hover {
  background-color: rgb(218, 55, 55);
}

.upload-section {
  height: 140px;
  display: flex;
  flex-direction: column;
  justify-content: left;
  /* align-items: left; */
  margin-left: 190px;
  margin-top: 20px;
  /* text-align: center; */
}

.upload-section h3 {
  color: #1b2821;
  font-family: "Outfit";
  font-size: 25px;
}

.upload-button {
  color: #fafefc;
  font-family: "Outfit";
  font-weight: 600;
  font-size: 15px;
  background-color: #4c956c;
  margin-top: 20px;
  width: 155px;
  height: 35px;
  border-radius: 8px;
  border: none;
}

.upload-button:hover{
  background-color: #3b7955;
}

.organizer {
  /* margin-left: 150px; Adjust this value to match the sidebar width */
  height: 100px;
  display: flex;
  justify-content: center;
  /* align-items: center; */
  margin-top: 50px;
  text-align: center;
}

.organizer h1 {
  color: #1b2821;
  font-family: "Outfit";
  font-size: 45px;
}
</style>
