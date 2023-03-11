import React from 'react';
import SmartLifeCharacter from 'components/ui/SmartLifeCharacter';
import Version from 'components/features/version/Main';
import Introduce from 'components/features/introduce/Main';

const Home = () =>  {
  return <div>
    <SmartLifeCharacter/>
    <Introduce/>
    <Version/>
  </div>
}

export default Home
