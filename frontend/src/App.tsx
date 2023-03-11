import React from 'react';
import {
  RouterProvider,
} from "react-router-dom";

import 'App.css';
import "index.css";
import router from 'Rotuer';

// TODO: https://zenn.dev/nakashi94/articles/f67fa9b54437da
function App() {
  return (
    <div className="app-general">
      <header className="app-header">
        <div>洋ナシオープンデータ</div>
      </header>
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
