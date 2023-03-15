import React from 'react';

import pearExample1 from 'images/example_1.png';
import pearExample2 from 'images/example_2.png';
import 'components/features/introduce/Main.css'
import { Box, Container, Typography } from '@mui/material';
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

        <Typography variant="subtitle1" align="left" color="text.secondary" paragraph>
          　新潟県の特産品の一つである高級洋ナシ「ル レクチエ」は、形状や外観などに関する出荷規格に基づいて等級判定が行われています。
          現在、多くの場合この等級判定は生産者の目視により行われており、作業負担と判定における個人差の発生が課題となっています。
          新潟大学工学部山﨑研究室では、これまでル レクチエの外観品質の自動評価の研究を進め、特に深層学習を利用した外観汚損検出から等級判定まで行うシステムを提案しており、その過程でル レクチエの外観画像のデータセットを構築してきました。
          <br/>
          　農業分野での深層学習技術の適用が進められては来ているものの、データセット構築に専門家の知識が必要であることや作物ごとに学習に適した量のデータが必要であること等が、実用化へのボトルネックになっています。
          この課題の克服へ貢献するために、当研究室では構築したル レクチエのデータセットを公開し、農業分野での深層学習の実用化がより進展することを期待します。
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
      <Typography variant="subtitle1" align="left" color="text.secondary" paragraph>
        <ol>
          <li>
            1. 新潟県名産品のル レクチエの外観品質を定義したデータの公開
          </li>
          <li>
            2. 外観品質を定義した独自フォーマットと物体検出時の一般的なアノテーションフォーマット(COCO)により提供
          </li>
          <li>
            3. 1000枚を超えるアノテーション済み画像データの提供
          </li>
        </ol>
      </Typography>

      <ul>
        <li><img src={pearExample1} className="pear-example-item" alt="pear_example_1" /></li>
        <li><img src={pearExample2} className="pear-example-item" alt="pear_example_2" /></li>
      </ul>
    </Container>
  </div>
}

export default Introduce
