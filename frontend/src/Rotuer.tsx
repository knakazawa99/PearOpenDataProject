import React from "react";
import {
  createBrowserRouter,
} from "react-router-dom";

import Home from 'components/pages/Home';
import AdminAuth from './components/pages/AdminAuth';

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
    element: AdminAuth(),
  }
]);

export default router
