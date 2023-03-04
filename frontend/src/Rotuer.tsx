import React from "react";
import {
  createBrowserRouter,
} from "react-router-dom";

import Home from 'components/pages/Home';

const router = createBrowserRouter([
  {
    path: "/",
    element: Home(),
  },
  {
    path: "/versions",
    element: <div>Versions!</div>,
  },
]);

// export const navigate = useNavigate();

export default router
