import { $t } from "@/plugins/i18n";
import { home } from "@/router/enums";
const { VITE_HIDE_HOME } = import.meta.env;
const Layout = () => import("@/layout/index.vue");

export default {
  path: "/",
  name: "Home",
  component: Layout,
  redirect: "/welcome",
  meta: {
    icon: "homeFilled",
    title: $t("menus.hshome"),
    rank: home
  },
  children: [
    {
      path: "/welcome",
      name: "Welcome",
      component: () => import("@/views/welcome/index.vue"),
      meta: {
        title: $t("menus.hshome"),
        showLink: VITE_HIDE_HOME === "true" ? false : true
      }
    },
    {
      path: "/error/403",
      name: "403",
      component: () => import("@/views/error/403.vue"),
      meta: {
        showLink: false,
        title: $t("menus.hsfourZeroOne")
      }
    },
    {
      path: "/error/404",
      name: "404",
      component: () => import("@/views/error/404.vue"),
      meta: {
        showLink: false,
        title: $t("menus.hsfourZeroFour")
      }
    },
    {
      path: "/error/500",
      name: "500",
      component: () => import("@/views/error/500.vue"),
      meta: {
        showLink: false,
        title: $t("menus.hsFive")
      }
    }
  ]
} satisfies RouteConfigsTable;
