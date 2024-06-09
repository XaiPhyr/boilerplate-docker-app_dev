import { createRouter, createWebHistory } from 'vue-router';
// import defineAbility from '@/utils/ability';

import DashboardPage from '../pages/dashboard/index.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: DashboardPage,
    },
  ],
});

router.beforeEach((to, from, next) => {
  console.log('TO: ', to);
  console.log('FROM: ', from);

  return next();
});

export default router;
