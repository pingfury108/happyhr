import { createRouter, createWebHashHistory } from "vue-router"
import AdminEmployees from "./components/AdminEmployees.vue"

const routes = [
  { path: '/employees', component: AdminEmployees },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes
})

export default router