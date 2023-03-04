import React from 'react';

import smartLifeCharacter from 'images/smart_life.gif';
import "components/ui/SmartLifeCharacter.css"

function SmartLifeCharacter() {
  return <img src={smartLifeCharacter} className="smart-life-logo" alt="logo" />
}

export default SmartLifeCharacter