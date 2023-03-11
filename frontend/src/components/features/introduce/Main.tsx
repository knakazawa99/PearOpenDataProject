import React from 'react';

import pearExample1 from 'images/example_1.png';
import pearExample2 from 'images/example_2.png';
import 'components/features/introduce/Main.css'
import { Box, Container, Link, Typography } from '@mui/material';
import SmartLifeCharacter from '../../ui/SmartLifeCharacter';

const Introduce = () => {
  return <div>
    <Box
      sx={{
        bgcolor: 'background.paper',
        pt: 6,
        pb: 6,
      }}
    >
      <Container maxWidth="md">
        <Typography
          component="h2"
          variant="h3"
          align="center"
          color="text.primary"
          gutterBottom
        >
          洋ナシオープンデータ
        </Typography>

        <SmartLifeCharacter/>

        <Typography variant="subtitle1" align="center" color="text.secondary" paragraph>
          新潟県名産のル レクチエはお歳暮等での需要から外観の品質が重要な果実です。
          そのことから外観の品質を定義した等級が外観の観点ごとに定められています。
          外観の観点は主に外観汚損(病害や傷)や形状により定められており、これまでこれらの評価は生産者の目視により行われていました。
          我々はこれまでの研究でル レクチエの外観品質を自動で評価するために深層学習を利用した手法を提案しており、その過程でル レクチエのデータセットを構築してきました。
          <br/>
          農業分野での深層学習技術の適用は期待されてはいるもののデータセットに専門家の知識が必要であることや必要なデータの量が多いため、様々な研究でデータセットの不足が課題になっています。
          この課題に少しでも貢献するため我々の研究室では、構築したル レクチエのデータセットを公開することにしました。
        </Typography>
      </Container>

    </Box>

    <Container maxWidth="md">
      <Typography
        component="h4"
        variant="h5"
        align="center"
        color="text.primary"
        gutterBottom
      >
        公開データの特徴
      </Typography>
      <Typography variant="subtitle1" align="center" color="text.secondary" paragraph>
        <div>
          新潟県名産品のル レクチエの外観品質を定義したデータの公開
        </div>
        <div>
          外観品質を定義した独自フォーマットと物体検出時の一般的なアノテーションフォーマット(COCO)により提供
        </div>
      </Typography>
      <ul>
        <li><img src={pearExample1} className="pear-example-item" alt="pear_example_1" /></li>
        <li><img src={pearExample2} className="pear-example-item" alt="pear_example_2" /></li>
      </ul>
    </Container>
  </div>
}

export default Introduce
