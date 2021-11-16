
import { concat } from "uint8arrays"
export function uint8chunksToBase64(chunks:any):string {
  const u8 = concat(chunks)
  var len = u8.byteLength;
  var binary = '';
  for (var i = 0; i < len; i++) {
    binary += String.fromCharCode(u8[i]);
  }
  return btoa(binary)
}