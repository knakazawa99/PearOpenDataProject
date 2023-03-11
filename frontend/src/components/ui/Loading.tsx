import React from 'react';

import smartLifeCharacter from 'images/smart_life_2.gif';
import "components/ui/Loading.css"

const Loading = ({ inverted = true, content="Loading..."}) => {
  return <div className="loading-body">
    <img src={smartLifeCharacter} className="loading-logo" alt="logo" />
    <div className="loading-text">{content}</div>
  </div>
}

export default Loading

