(window.webpackJsonp=window.webpackJsonp||[]).push([[4],{371:function(t,e,l){var content=l(373);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[t.i,content,""]]),content.locals&&(t.exports=content.locals);(0,l(79).default)("31aeb2da",content,!0,{sourceMap:!1})},372:function(t,e,l){"use strict";l(371)},373:function(t,e,l){var c=l(78)(!1);c.push([t.i,".rect[data-v-2d68292e]{border:1.2px solid;display:table-cell;text-align:center;vertical-align:middle}.circle[data-v-2d68292e]{width:40px;height:40px;line-height:40px;font-size:30px}.circle[data-v-2d68292e],.circle--small[data-v-2d68292e]{display:inline-block;border:1.2px solid;border-radius:50%;text-align:center}.circle--small[data-v-2d68292e]{width:20px;height:20px;line-height:16px;margin-top:15px}.circle--dakuten[data-v-2d68292e]{width:35px;height:35px;line-height:35px;margin-top:5px}.circle--dakuten[data-v-2d68292e],.dakuten[data-v-2d68292e]{display:inline-block;border:1.2px solid;border-radius:50%;text-align:center}.dakuten[data-v-2d68292e]{width:10px;height:10px;line-height:20px}.opened[data-v-2d68292e]{background-color:grey}.hide[data-v-2d68292e]{background-color:#000}",""]),t.exports=c},374:function(t,e,l){"use strict";l.r(e);l(114);var c=l(47),n=Object(c.a)({name:"SquareCom",props:{charaInfo:{type:Array,default:null}},setup:function(t){return{props:t,isSmall:function(t){return["ゃ","ゅ","ょ","ぁ","ぃ","ぅ","ぇ","ぉ","っ","ゎ"].includes(t)},isDakuon:function(t){return["が","ぎ","ぐ","げ","ご","ざ","じ","ず","ぜ","ぞ","だ","ぢ","づ","で","ど","ば","び","ぶ","べ","ぼ"].includes(t)},isHanDakuon:function(t){return["ぱ","ぴ","ぷ","ぺ","ぽ"].includes(t)},changeChar:{"ゃ":"や","ゅ":"ゆ","ょ":"よ","ぁ":"あ","ぃ":"い","ぅ":"う","ぇ":"え","ぉ":"お","っ":"つ","ゎ":"わ","が":"か","ぎ":"き","ぐ":"く","げ":"け","ご":"こ","ざ":"さ","じ":"し","ず":"す","ぜ":"せ","ぞ":"そ","だ":"た","ぢ":"ち","づ":"つ","で":"て","ど":"と","ば":"は","び":"ひ","ぶ":"ふ","べ":"へ","ぼ":"ほ","ぱ":"は","ぴ":"ひ","ぷ":"ふ","ぺ":"へ","ぽ":"ほ"}}}}),d=(l(372),l(76)),component=Object(d.a)(n,(function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("div",[l("div",{staticClass:"d-flex"},t._l(t.props.charaInfo,(function(e,c){return l("div",{key:c,staticClass:"mx-2"},[e.isHide?l("div",[t.isSmall(e.char)?l("div",{staticClass:"circle--small hide"},[t._v("・")]):t.isDakuon(e.char)||t.isHanDakuon(e.char)?l("div",{staticStyle:{position:"relative"}},[l("div",{staticClass:"circle hide"},[t._v("・")]),t._v(" "),l("div",{staticClass:"dakuten hide",staticStyle:{position:"absolute",top:"0"}})]):l("div",{staticClass:"circle hide"},[t._v("・")])]):t.isSmall(e.char)?l("div",{staticClass:"circle--small"},[t._v(t._s(t.changeChar[e.char]))]):t.isDakuon(e.char)?l("div",{staticStyle:{position:"relative"}},[l("div",{staticClass:"circle"},[t._v(t._s(t.changeChar[e.char]))]),t._v(" "),l("div",{staticClass:"dakuten",staticStyle:{position:"absolute",top:"0"}},[t._v("゛")])]):t.isHanDakuon(e.char)?l("div",{staticStyle:{position:"relative"}},[l("div",{staticClass:"circle"},[t._v(t._s(t.changeChar[e.char]))]),t._v(" "),l("div",{staticClass:"dakuten",staticStyle:{position:"absolute",top:"0"}},[t._v("゜")])]):l("div",{staticClass:"circle"},[t._v("\n        "+t._s(e.char)+"\n      ")])])})),0)])}),[],!1,null,"2d68292e",null);e.default=component.exports}}]);