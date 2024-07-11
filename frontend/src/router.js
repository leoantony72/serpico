import { createRouter, createWebHistory } from 'vue-router';
import Organizer from './components/Organize.vue';
import Trash from './components/Trash.vue';

const routes = [
  { path: '/', component: Organizer },
  { path: '/organizer', component: Organizer },
  // { path: '/trash', component: Trash },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
