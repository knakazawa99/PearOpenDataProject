import React from 'react';

import pearExample1 from 'images/example_1.png';
import pearExample2 from 'images/example_2.png';
import 'components/features/introduce/Main.css'

const Introduce = () => {
  return <div>
    <div>
      新潟県名産品のル レクチエの外観品質を定義したデータの公開
    </div>
    <div>
      一般的なアノテーションフォーマット(COCO)により提供
    </div>
    <ul>
      <li><img src={pearExample1} className="pear-example-item" alt="pear_example_1" /></li>
      <li><img src={pearExample2} className="pear-example-item" alt="pear_example_2" /></li>
    </ul>
  </div>
}

export default Introduce
