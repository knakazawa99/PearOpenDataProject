import * as React from 'react';
import { Box, Typography } from '@mui/material';


export default function Term() {
  return <div>

    <div>
      <Typography component="h5" variant="h5">
        利用規約
      </Typography>
    </div>
    <Box
      sx={{
        height: 200,
        overflow: "hidden",
        overflowY: "scroll",
      }}
      >
      <Typography paragraph>
        この利用規約（以下、「本規約」といいます。）は、新潟大学山﨑達也研究室（以下、「当研究室」）がウェブサイト上で提供するサービス「PearOpenData」（以下、「本サービス」）の提供条件と、
        本サービスを利用するユーザーに遵守していただくべき事項、運営元とユーザーの間の権利義務関係が定められています。
      </Typography>

      <div>
        <Typography component="h5" variant="subtitle1">第1条（適用）</Typography>
        <Typography paragraph>
          <ol>
            <li>本規約は、本サービスの提供条件と本サービスの利用に関する当研究室とユーザーとの間の権利義務関係を定めることを目的として、ユーザーと当研究室との間の本サービスの利用に関わるすべての関係に適用されます。</li>
            <li>本規約のほか、当研究室が定め、本サービス上に掲載する本サービスの利用に関するルール、各種規定は、本規約の一部を構成するものとします。</li>
            <li>本規約の内容と、前項のルール、各種規定、その他の本規約外における本サービスの説明とが異なる場合は、本規約の規定が優先して適用されるものとします。</li>
          </ol>
        </Typography>
      </div>

    <div>
      <Typography component="h5" variant="subtitle1">第2条（禁止事項）</Typography>
      <Typography paragraph>
        ユーザーは、本サービスの利用にあたり、以下の行為をしてはなりません。
        <ol>
          <li>法令または公序良俗に違反する行為</li>
          <li>犯罪行為に関連する行為</li>
          <li>本サービスの内容等、本サービスに含まれる著作権、商標権ほか知的財産権を侵害する行為</li>
          <li>当研究室、ほかのユーザー、またはその他第三者のサーバーまたはネットワークの機能を破壊したり、妨害したりする行為</li>
          <li>本サービスによって得られた情報を商業的に利用する行為</li>
          <li>当研究室のサービスの運営を妨害するおそれのある行為</li>
          <li>不正アクセスをし、またはこれを試みる行為</li>
          <li>他のユーザーに関する個人情報等を収集または蓄積する行為</li>
          <li>不正な目的を持って本サービスを利用する行為</li>
          <li>本サービスの他のユーザーまたはその他の第三者に不利益、損害、不快感を与える行為</li>
          <li>他のユーザーに成りすます行為</li>
          <li>当研究室が許諾しない本サービス上での宣伝、広告、勧誘、または営業行為</li>
          <li>当研究室のサービスに関連して、反社会的勢力に対して直接または間接に利益を供与する行為</li>
          <li>その他、当研究室が不適切と判断する行為</li>
        </ol>
      </Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第3条（本サービスの提供の停止等）</Typography>
      <Typography paragraph>
        当研究室は、以下のいずれかの事由があると判断した場合、ユーザーに事前に通知することなく本サービスの全部または一部の提供を停止または中断することができるものとします。
        <ol>
          <li>本サービスにかかるコンピュータシステムの保守点検または更新を行う場合</li>
          <li>地震、落雷、火災、停電または天災などの不可抗力により、本サービスの提供が困難となった場合</li>
          <li>コンピュータまたは通信回線等が事故により停止した場合</li>
          <li>その他、当研究室が本サービスの提供が困難と判断した場合</li>
        </ol>
        当研究室は、サービスの提供の停止または中断により、ユーザーまたは第三者が被ったいかなる不利益または損害についても、一切の責任を負わないものとします。
      </Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第4条（保証の否認および免責事項）</Typography>
      <Typography paragraph>
        当研究室は、本サービスに事実上または法律上の瑕疵（安全性、信頼性、正確性、完全性、有効性、特定の目的への適合性、セキュリティなどに関する欠陥、エラーやバグ、権利侵害などを含みます。）がないことを明示的にも黙示的にも保証しておりません。<br/>
        当研究室は、本サービスに起因してユーザーに生じたあらゆる損害について一切の責任を負いません。<br/>
        前項ただし書に定める場合であっても、当研究室は、当研究室の過失（重過失を除きます。）による債務不履行または不法行為によりユーザーに生じた損害のうち特別な事情から生じた損害（当研究室またはユーザーが損害発生につき予見し、または予見し得た場合を含みます。）について一切の責任を負いません。<br/>
        また、当研究室の過失（重過失を除きます。）による債務不履行または不法行為によりユーザーに生じた損害の賠償は、ユーザーから当該損害が発生した月に受領した利用料の額を上限とします。<br/>
        当研究室は、本サービスに関して、ユーザーと他のユーザーまたは第三者との間において生じた取引、連絡または紛争等について一切責任を負いません。
      </Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第5条（サービス内容の変更等）</Typography>
      <Typography paragraph>
        当研究室は、ユーザーに通知することなく、本サービスの内容を変更しまたは本サービスの提供を中止することができるものとし、これによってユーザーに生じた損害について一切の責任を負いません。
      </Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第6条（利用規約の変更）</Typography>
      <Typography paragraph>
        当研究室は、必要と判断した場合には、ユーザーに通知することなくいつでも本規約を変更することができるものとします。<br/>
        なお、本規約の変更後、本サービスの利用を開始した場合には、当該ユーザーは変更後の規約に同意したものとみなします。
      </Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">
        第7条（通知または連絡）
      </Typography>
      <Typography paragraph>
        ユーザーと当研究室との間の通知または連絡は、当研究室の定める方法によって行うものとします。<br/>
        当研究室は、ユーザーから、当研究室が別途定める方式に従った変更届け出がない限り、現在登録されている連絡先が有効なものとみなして当該連絡先へ通知または連絡を行い、これらは、発信時にユーザーへ到達したものとみなします。
      </Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第8条（個人情報の取り扱い）</Typography>
      <Typography paragraph></Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第9条（使用許諾）</Typography>
      <Typography paragraph></Typography>
    </div>

    <div>
      <Typography component="h5" variant="subtitle1">第10条（準拠法・裁判管轄）</Typography>
      <Typography paragraph>
        本規約の解釈にあたっては、日本法を準拠法とします。
        本サービスに関して紛争が生じた場合には、当研究室の本店所在地を管轄する裁判所を専属的合意管轄とします。
      </Typography>
    </div>
  </Box></div>
}