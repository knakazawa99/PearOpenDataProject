import React from 'react';
import {
  RouterProvider,
} from "react-router-dom";

import Copyright from 'components/ui/Copyright';
import router from 'Rotuer';

import 'App.css';
import "index.css";

function App() {
  return (
    <div className="app-general">
      <RouterProvider router={router} />
      <Copyright/>
    </div>
  );
}

export default App;
