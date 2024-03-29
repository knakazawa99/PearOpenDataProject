import React from "react";
import {
  createBrowserRouter,
} from "react-router-dom";

import Home from 'components/pages/Home';
import AdminAuthPage from 'components/pages/AdminAuth';
import AdminVersionPage from 'components/pages/AdminVersion';
import AdminPage from 'components/pages/Admin';

const router = createBrowserRouter([
  {
    path: "/",
    element: Home(),
  },
  {
    path: "/versions",
    element: <div>Versions!</div>,
  },
  {
    path: "/admin/signin",
    element: AdminAuthPage(),
  },
  {
    path: "/admin/versions",
    element: AdminVersionPage(),
  },
  {
    path: "/admin/admins",
    element: AdminPage(),
  },
]);

export default router
