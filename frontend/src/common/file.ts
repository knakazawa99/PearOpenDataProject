export const getFileNameFromContentDisposition = (contentDisposition: any):string => {
  const regex = /filename="(.*)"/; // ファイル名を抽出する正規表現
  const match = regex.exec(contentDisposition);
  const fileName = match?.[1]; // マッチした文字列の2番目の要素がファイル名

  if (!fileName) {
    return ""
  }
  return fileName;
}