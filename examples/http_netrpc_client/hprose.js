// Hprose for JavaScript v2.0.16
// Copyright (c) 2008-2016 http://hprose.com
// Hprose is freely distributable under the MIT license.
// For all details and documentation:
// https://github.com/hprose/hprose-js


(function(global){'use strict';global.hprose={};})(this||[eval][0]('this'));(function(global){'use strict';if(typeof global.setTimeout!=="undefined"){return;}
if(typeof global.require!=="function"){return;}
var deviceone;try{deviceone=global.require("deviceone");}
catch(e){return;}
if(!deviceone){return;}
global.setTimeout=function(func,delay){if(delay<=0){delay=1;}
var timer=deviceone.mm("do_Timer");timer.delay=delay;timer.interval=delay;timer.on('tick',function(){timer.stop();func();});timer.start();return timer;};global.clearTimeout=function(timer){if(timer.isStart()){timer.stop();}};})(this||[eval][0]('this'));(function(global){'use strict';if(typeof(global.btoa)==="undefined"){global.btoa=(function(){var base64EncodeChars='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'.split('');return function(str){var buf,i,j,len,r,l,c;i=j=0;len=str.length;r=len%3;len=len-r;l=(len/3)<<2;if(r>0){l+=4;}
buf=new Array(l);while(i<len){c=str.charCodeAt(i++)<<16|str.charCodeAt(i++)<<8|str.charCodeAt(i++);buf[j++]=base64EncodeChars[c>>18]+
base64EncodeChars[c>>12&0x3f]+
base64EncodeChars[c>>6&0x3f]+
base64EncodeChars[c&0x3f];}
if(r===1){c=str.charCodeAt(i++);buf[j++]=base64EncodeChars[c>>2]+
base64EncodeChars[(c&0x03)<<4]+"==";}
else if(r===2){c=str.charCodeAt(i++)<<8|str.charCodeAt(i++);buf[j++]=base64EncodeChars[c>>10]+
base64EncodeChars[c>>4&0x3f]+
base64EncodeChars[(c&0x0f)<<2]+"=";}
return buf.join('');};})();}
if(typeof(global.atob)==="undefined"){global.atob=(function(){var base64DecodeChars=[-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,62,-1,-1,-1,63,52,53,54,55,56,57,58,59,60,61,-1,-1,-1,-1,-1,-1,-1,0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,-1,-1,-1,-1,-1,-1,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,-1,-1,-1,-1,-1];return function(str){var c1,c2,c3,c4;var i,j,len,r,l,out;len=str.length;if(len%4!==0){return'';}
if(/[^ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789\+\/\=]/.test(str)){return'';}
if(str.charAt(len-2)==='='){r=1;}
else if(str.charAt(len-1)==='='){r=2;}
else{r=0;}
l=len;if(r>0){l-=4;}
l=(l>>2)*3+r;out=new Array(l);i=j=0;while(i<len){c1=base64DecodeChars[str.charCodeAt(i++)];if(c1===-1){break;}
c2=base64DecodeChars[str.charCodeAt(i++)];if(c2===-1){break;}
out[j++]=String.fromCharCode((c1<<2)|((c2&0x30)>>4));c3=base64DecodeChars[str.charCodeAt(i++)];if(c3===-1){break;}
out[j++]=String.fromCharCode(((c2&0x0f)<<4)|((c3&0x3c)>>2));c4=base64DecodeChars[str.charCodeAt(i++)];if(c4===-1){break;}
out[j++]=String.fromCharCode(((c3&0x03)<<6)|c4);}
return out.join('');};})();}})(this||[eval][0]('this'));(function(global,undefined){'use strict';var propertyToMethod=function(prop){if('get'in prop&&'set'in prop){return function(){if(arguments.length===0){return prop.get();}
prop.set(arguments[0]);};}
if('get'in prop){return prop.get;}
if('set'in prop){return prop.set;}}
var defineProperties=(typeof Object.defineProperties!=='function'?function(obj,properties){var buildinMethod=['toString','toLocaleString','valueOf','hasOwnProperty','isPrototypeOf','propertyIsEnumerable','constructor'];buildinMethod.forEach(function(name){var prop=properties[name];if('value'in prop){obj[name]=prop.value;}});for(var name in properties){var prop=properties[name];obj[name]=undefined;if('value'in prop){obj[name]=prop.value;}
else if('get'in prop||'set'in prop){obj[name]=propertyToMethod(prop);}}}:function(obj,properties){for(var name in properties){var prop=properties[name];if('get'in prop||'set'in prop){properties[name]={value:propertyToMethod(prop)};}}
Object.defineProperties(obj,properties);});var Temp=function(){};var createObject=(typeof Object.create!=='function'?function(prototype,properties){if(typeof prototype!='object'&&typeof prototype!='function'){throw new TypeError('prototype must be an object or function');}
Temp.prototype=prototype;var result=new Temp();Temp.prototype=null;if(properties){defineProperties(result,properties);}
return result;}:function(prototype,properties){if(properties){for(var name in properties){var prop=properties[name];if('get'in prop||'set'in prop){properties[name]={value:propertyToMethod(prop)};}}
return Object.create(prototype,properties);}
return Object.create(prototype);});var generic=function(method){if(typeof method!=="function"){throw new TypeError(method+" is not a function");}
return function(context){return method.apply(context,Array.prototype.slice.call(arguments,1));};}
var toArray=function(arrayLikeObject){var n=arrayLikeObject.length;var a=new Array(n);for(var i=0;i<n;++i){a[i]=arrayLikeObject[i];}
return a;};var toBinaryString=function(bytes){if(bytes instanceof ArrayBuffer){bytes=new Uint8Array(bytes);}
var n=bytes.length;if(n<0xFFFF){return String.fromCharCode.apply(String,toArray(bytes));}
var remain=n&0x7FFF;var count=n>>15;var a=new Array(remain?count+1:count);for(var i=0;i<count;++i){a[i]=String.fromCharCode.apply(String,toArray(bytes.subarray(i<<15,(i+1)<<15)));}
if(remain){a[count]=String.fromCharCode.apply(String,toArray(bytes.subarray(count<<15,n)));}
return a.join('');};var toUint8Array=function(bs){var n=bs.length;var data=new Uint8Array(n);for(var i=0;i<n;i++){data[i]=bs.charCodeAt(i)&0xFF;}
return data;};var parseuri=function(url){var pattern=new RegExp("^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?");var matches=url.match(pattern);var host=matches[4].split(':',2);return{protocol:matches[1],host:matches[4],hostname:host[0],port:parseInt(host[1],10)||0,path:matches[5],query:matches[7],fragment:matches[9]};}
var isObjectEmpty=function(obj){if(obj){var prop;for(prop in obj){return false;}}
return true;}
global.hprose.defineProperties=defineProperties;global.hprose.createObject=createObject;global.hprose.generic=generic;global.hprose.toBinaryString=toBinaryString;global.hprose.toUint8Array=toUint8Array;global.hprose.toArray=toArray;global.hprose.parseuri=parseuri;global.hprose.isObjectEmpty=isObjectEmpty;})(this||[eval][0]('this'));(function(global,undefined){'use strict';if(!Function.prototype.bind){Function.prototype.bind=function(oThis){if(typeof this!=='function'){throw new TypeError('Function.prototype.bind - what is trying to be bound is not callable');}
var aArgs=Array.prototype.slice.call(arguments,1),toBind=this,NOP=function(){},bound=function(){return toBind.apply(this instanceof NOP?this:oThis,aArgs.concat(Array.prototype.slice.call(arguments)));};if(this.prototype){NOP.prototype=this.prototype;}
bound.prototype=new NOP();return bound;};}
if(!Array.prototype.indexOf){Array.prototype.indexOf=function(searchElement){if(this===null||this===undefined){throw new TypeError('"this" is null or not defined');}
var o=Object(this);var len=o.length>>>0;if(len===0){return-1;}
var n=+Number(arguments[1])||0;if(Math.abs(n)===Infinity){n=0;}
if(n>=len){return-1;}
var k=Math.max(n>=0?n:len-Math.abs(n),0);while(k<len){if(k in o&&o[k]===searchElement){return k;}
k++;}
return-1;};}
if(!Array.prototype.lastIndexOf){Array.prototype.lastIndexOf=function(searchElement){if(this===null||this===undefined){throw new TypeError('"this" is null or not defined');}
var o=Object(this);var len=o.length>>>0;if(len===0){return-1;}
var n=+Number(arguments[1])||0;if(Math.abs(n)===Infinity){n=0;}
for(var k=n>=0?Math.min(n,len-1):len-Math.abs(n);k>=0;k--){if(k in o&&o[k]===searchElement){return k;}}
return-1;};}
if(!Array.prototype.filter){Array.prototype.filter=function(fun){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof fun!=="function"){throw new TypeError(fun+" is not a function");}
var res=[];var thisp=arguments[1];for(var i=0;i<len;i++){if(i in t){var val=t[i];if(fun.call(thisp,val,i,t)){res.push(val);}}}
return res;};}
if(!Array.prototype.forEach){Array.prototype.forEach=function(fun){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof fun!=="function"){throw new TypeError(fun+" is not a function");}
var thisp=arguments[1];for(var i=0;i<len;i++){if(i in t){fun.call(thisp,t[i],i,t);}}};}
if(!Array.prototype.every){Array.prototype.every=function(fun){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof fun!=="function"){throw new TypeError(fun+" is not a function");}
var thisp=arguments[1];for(var i=0;i<len;i++){if(i in t&&!fun.call(thisp,t[i],i,t)){return false;}}
return true;};}
if(!Array.prototype.map){Array.prototype.map=function(fun){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof fun!=="function"){throw new TypeError(fun+" is not a function");}
var thisp=arguments[1];var res=new Array(len);for(var i=0;i<len;i++){if(i in t){res[i]=fun.call(thisp,t[i],i,t);}}
return res;};}
if(!Array.prototype.some){Array.prototype.some=function(fun){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof fun!=="function"){throw new TypeError(fun+" is not a function");}
var thisp=arguments[1];for(var i=0;i<len;i++){if(i in t&&fun.call(thisp,t[i],i,t)){return true;}}
return false;};}
if(!Array.prototype.reduce){Array.prototype.reduce=function(callbackfn){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof callbackfn!=="function"){throw new TypeError("First argument is not callable");}
if(len===0&&arguments.length===1){throw new TypeError("Array length is 0 and no second argument");}
var i=0,accumulator;if(arguments.length>=2){accumulator=arguments[1];}
else{accumulator=t[0];i=1;}
for(;i<len;++i){if(i in t){accumulator=callbackfn.call(undefined,accumulator,t[i],i,t);}}
return accumulator;};}
if(!Array.prototype.reduceRight){Array.prototype.reduceRight=function(callbackfn){if(this===null||this===undefined){throw new TypeError("this is null or not defined");}
var t=Object(this);var len=t.length>>>0;if(typeof callbackfn!=="function"){throw new TypeError("First argument is not callable");}
if(len===0&&arguments.length===1){throw new TypeError("Array length is 0 and no second argument");}
var k=len-1;var accumulator;if(arguments.length>=2){accumulator=arguments[1];}
else{do{if(k in t){accumulator=t[k--];break;}
if(--k<0){throw new TypeError("Array contains no values");}}while(true);}
while(k>=0){if(k in t){accumulator=callbackfn.call(undefined,accumulator,t[k],k,t);}
k--;}
return accumulator;};}
if(!Array.prototype.includes){Array.prototype.includes=function(searchElement){var O=Object(this);var len=parseInt(O.length,10)||0;if(len===0){return false;}
var n=parseInt(arguments[1],10)||0;var k;if(n>=0){k=n;}
else{k=len+n;if(k<0){k=0;}}
var currentElement;while(k<len){currentElement=O[k];if(searchElement===currentElement||(searchElement!==searchElement&&currentElement!==currentElement)){return true;}
k++;}
return false;};}
if(!Array.prototype.find){Array.prototype.find=function(predicate){if(this===null||this===undefined){throw new TypeError('Array.prototype.find called on null or undefined');}
if(typeof predicate!=='function'){throw new TypeError('predicate must be a function');}
var list=Object(this);var length=list.length>>>0;var thisArg=arguments[1];var value;for(var i=0;i<length;i++){value=list[i];if(predicate.call(thisArg,value,i,list)){return value;}}
return undefined;};}
if(!Array.prototype.findIndex){Array.prototype.findIndex=function(predicate){if(this===null||this===undefined){throw new TypeError('Array.prototype.findIndex called on null or undefined');}
if(typeof predicate!=='function'){throw new TypeError('predicate must be a function');}
var list=Object(this);var length=list.length>>>0;var thisArg=arguments[1];var value;for(var i=0;i<length;i++){value=list[i];if(predicate.call(thisArg,value,i,list)){return i;}}
return-1;};}
if(!Array.prototype.fill){Array.prototype.fill=function(value){if(this===null||this===undefined){throw new TypeError('this is null or not defined');}
var O=Object(this);var len=O.length>>>0;var start=arguments[1];var relativeStart=start>>0;var k=relativeStart<0?Math.max(len+relativeStart,0):Math.min(relativeStart,len);var end=arguments[2];var relativeEnd=end===undefined?len:end>>0;var f=relativeEnd<0?Math.max(len+relativeEnd,0):Math.min(relativeEnd,len);while(k<f){O[k]=value;k++;}
return O;};}
if(!Array.prototype.copyWithin){Array.prototype.copyWithin=function(target,start){if(this===null||this===undefined){throw new TypeError('this is null or not defined');}
var O=Object(this);var len=O.length>>>0;var relativeTarget=target>>0;var to=relativeTarget<0?Math.max(len+relativeTarget,0):Math.min(relativeTarget,len);var relativeStart=start>>0;var from=relativeStart<0?Math.max(len+relativeStart,0):Math.min(relativeStart,len);var end=arguments[2];var relativeEnd=end===undefined?len:end>>0;var f=relativeEnd<0?Math.max(len+relativeEnd,0):Math.min(relativeEnd,len);var count=Math.min(f-from,len-to);var direction=1;if(from<to&&to<(from+count)){direction=-1;from+=count-1;to+=count-1;}
while(count>0){if(from in O){O[to]=O[from];}
else{delete O[to];}
from+=direction;to+=direction;count--;}
return O;};}
if(!Array.isArray){Array.isArray=function(arg){return Object.prototype.toString.call(arg)==='[object Array]';};}
if(!Array.from){Array.from=(function(){var toStr=Object.prototype.toString;var isCallable=function(fn){return typeof fn==='function'||toStr.call(fn)==='[object Function]';};var toInteger=function(value){var number=Number(value);if(isNaN(number)){return 0;}
if(number===0||!isFinite(number)){return number;}
return(number>0?1:-1)*Math.floor(Math.abs(number));};var maxSafeInteger=Math.pow(2,53)-1;var toLength=function(value){var len=toInteger(value);return Math.min(Math.max(len,0),maxSafeInteger);};return function(arrayLike){var C=this;var items=Object(arrayLike);if(arrayLike===null||arrayLike===undefined){throw new TypeError("Array.from requires an array-like object - not null or undefined");}
var mapFn=arguments.length>1?arguments[1]:void undefined;var T;if(typeof mapFn!=='undefined'){if(!isCallable(mapFn)){throw new TypeError('Array.from: when provided, the second argument must be a function');}
if(arguments.length>2){T=arguments[2];}}
var len=toLength(items.length);var A=isCallable(C)?Object(new C(len)):new Array(len);var k=0;var kValue;while(k<len){kValue=items[k];if(mapFn){A[k]=typeof T==='undefined'?mapFn(kValue,k):mapFn.call(T,kValue,k);}
else{A[k]=kValue;}
k+=1;}
A.length=len;return A;};}());}
if(!Array.of){Array.of=function(){return Array.prototype.slice.call(arguments);};}
if(!String.prototype.startsWith){String.prototype.startsWith=function(searchString,position){position=position||0;return this.substr(position,searchString.length)===searchString;};}
if(!String.prototype.endsWith){String.prototype.endsWith=function(searchString,position){var subjectString=this.toString();if(typeof position!=='number'||!isFinite(position)||Math.floor(position)!==position||position>subjectString.length){position=subjectString.length;}
position-=searchString.length;var lastIndex=subjectString.indexOf(searchString,position);return lastIndex!==-1&&lastIndex===position;};}
if(!String.prototype.includes){String.prototype.includes=function(){if(typeof arguments[1]==="number"){if(this.length<arguments[0].length+arguments[1].length){return false;}
else{return this.substr(arguments[1],arguments[0].length)===arguments[0];}}
else{return String.prototype.indexOf.apply(this,arguments)!==-1;}};}
if(!String.prototype.repeat){String.prototype.repeat=function(count){var str=this.toString();count=+count;if(count!==count){count=0;}
if(count<0){throw new RangeError('repeat count must be non-negative');}
if(count===Infinity){throw new RangeError('repeat count must be less than infinity');}
count=Math.floor(count);if(str.length===0||count===0){return'';}
if(str.length*count>=1<<28){throw new RangeError('repeat count must not overflow maximum string size');}
var rpt='';for(;;){if((count&1)===1){rpt+=str;}
count>>>=1;if(count===0){break;}
str+=str;}
return rpt;};}
if(!String.prototype.trim){String.prototype.trim=function(){return this.toString().replace(/^[\s\xa0]+|[\s\xa0]+$/g,'');};}
if(!String.prototype.trimLeft){String.prototype.trimLeft=function(){return this.toString().replace(/^[\s\xa0]+/,'');};}
if(!String.prototype.trimRight){String.prototype.trimRight=function(){return this.toString().replace(/[\s\xa0]+$/,'');};}
if(!Object.keys){Object.keys=(function(){var hasOwnProperty=Object.prototype.hasOwnProperty,hasDontEnumBug=!({toString:null}).propertyIsEnumerable('toString'),dontEnums=['toString','toLocaleString','valueOf','hasOwnProperty','isPrototypeOf','propertyIsEnumerable','constructor'],dontEnumsLength=dontEnums.length;return function(obj){if(typeof obj!=='object'&&typeof obj!=='function'||obj===null){throw new TypeError('Object.keys called on non-object');}
var result=[];for(var prop in obj){if(hasOwnProperty.call(obj,prop)){result.push(prop);}}
if(hasDontEnumBug){for(var i=0;i<dontEnumsLength;i++){if(hasOwnProperty.call(obj,dontEnums[i])){result.push(dontEnums[i]);}}}
return result;};})();}
if(!Date.now){Date.now=function(){return+(new Date());};}
if(!Date.prototype.toISOString){var f=function(n){return n<10?'0'+n:n;};Date.prototype.toISOString=function(){return this.getUTCFullYear()+'-'+
f(this.getUTCMonth()+1)+'-'+
f(this.getUTCDate())+'T'+
f(this.getUTCHours())+':'+
f(this.getUTCMinutes())+':'+
f(this.getUTCSeconds())+'Z';};}
var generic=global.hprose.generic;function genericMethods(obj,properties){var proto=obj.prototype;for(var i=0,len=properties.length;i<len;i++){var property=properties[i];var method=proto[property];if(typeof method==='function'&&typeof obj[property]==='undefined'){obj[property]=generic(method);}}}
genericMethods(Array,["pop","push","reverse","shift","sort","splice","unshift","concat","join","slice","indexOf","lastIndexOf","filter","forEach","every","map","some","reduce","reduceRight","includes","find","findIndex","fill","copyWithin"]);genericMethods(String,['quote','substring','toLowerCase','toUpperCase','charAt','charCodeAt','indexOf','lastIndexOf','include','startsWith','endsWith','repeat','trim','trimLeft','trimRight','toLocaleLowerCase','toLocaleUpperCase','match','search','replace','split','substr','concat','slice']);})(this||[eval][0]('this'));(function(global){'use strict';var hasWeakMap='WeakMap'in global;var hasMap='Map'in global;var hasForEach=true;if(hasMap){hasForEach='forEach'in new global.Map();}
if(hasWeakMap&&hasMap&&hasForEach){return;}
var hasObject_create='create'in Object;var createNPO=function(){return hasObject_create?Object.create(null):{};};var namespaces=createNPO();var count=0;var reDefineValueOf=function(obj){var privates=createNPO();var baseValueOf=obj.valueOf;var valueOf=function(namespace,n){if((this===obj)&&(n in namespaces)&&(namespaces[n]===namespace)){if(!(n in privates)){privates[n]=createNPO();}
return privates[n];}
else{return baseValueOf.apply(this,arguments);}};if(hasObject_create&&'defineProperty'in Object){Object.defineProperty(obj,'valueOf',{value:valueOf,writable:true,configurable:true,enumerable:false});}
else{obj.valueOf=valueOf;}};if(!hasWeakMap){global.WeakMap=function WeakMap(){var namespace=createNPO();var n=count++;namespaces[n]=namespace;var map=function(key){if(key!==Object(key)){throw new Error('value is not a non-null object');}
var privates=key.valueOf(namespace,n);if(privates!==key.valueOf()){return privates;}
reDefineValueOf(key);return key.valueOf(namespace,n);};var m=this;if(hasObject_create){m=Object.create(WeakMap.prototype,{get:{value:function(key){return map(key).value;},writable:false,configurable:false,enumerable:false},set:{value:function(key,value){map(key).value=value;},writable:false,configurable:false,enumerable:false},has:{value:function(key){return'value'in map(key);},writable:false,configurable:false,enumerable:false},'delete':{value:function(key){return delete map(key).value;},writable:false,configurable:false,enumerable:false},clear:{value:function(){delete namespaces[n];n=count++;namespaces[n]=namespace;},writable:false,configurable:false,enumerable:false}});}
else{m.get=function(key){return map(key).value;};m.set=function(key,value){map(key).value=value;};m.has=function(key){return'value'in map(key);};m['delete']=function(key){return delete map(key).value;};m.clear=function(){delete namespaces[n];n=count++;namespaces[n]=namespace;};}
if(arguments.length>0&&Array.isArray(arguments[0])){var iterable=arguments[0];for(var i=0,len=iterable.length;i<len;i++){m.set(iterable[i][0],iterable[i][1]);}}
return m;};}
if(!hasMap){var objectMap=function(){var namespace=createNPO();var n=count++;var nullMap=createNPO();namespaces[n]=namespace;var map=function(key){if(key===null){return nullMap;}
var privates=key.valueOf(namespace,n);if(privates!==key.valueOf()){return privates;}
reDefineValueOf(key);return key.valueOf(namespace,n);};return{get:function(key){return map(key).value;},set:function(key,value){map(key).value=value;},has:function(key){return'value'in map(key);},'delete':function(key){return delete map(key).value;},clear:function(){delete namespaces[n];n=count++;namespaces[n]=namespace;}};};var noKeyMap=function(){var map=createNPO();return{get:function(){return map.value;},set:function(_,value){map.value=value;},has:function(){return'value'in map;},'delete':function(){return delete map.value;},clear:function(){map=createNPO();}};};var scalarMap=function(){var map=createNPO();return{get:function(key){return map[key];},set:function(key,value){map[key]=value;},has:function(key){return key in map;},'delete':function(key){return delete map[key];},clear:function(){map=createNPO();}};};if(!hasObject_create){var stringMap=function(){var map={};return{get:function(key){return map['!'+key];},set:function(key,value){map['!'+key]=value;},has:function(key){return('!'+key)in map;},'delete':function(key){return delete map['!'+key];},clear:function(){map={};}};};}
global.Map=function Map(){var map={'number':scalarMap(),'string':hasObject_create?scalarMap():stringMap(),'boolean':scalarMap(),'object':objectMap(),'function':objectMap(),'unknown':objectMap(),'undefined':noKeyMap(),'null':noKeyMap()};var size=0;var keys=[];var m=this;if(hasObject_create){m=Object.create(Map.prototype,{size:{get:function(){return size;},configurable:false,enumerable:false},get:{value:function(key){return map[typeof(key)].get(key);},writable:false,configurable:false,enumerable:false},set:{value:function(key,value){if(!this.has(key)){keys.push(key);size++;}
map[typeof(key)].set(key,value);},writable:false,configurable:false,enumerable:false},has:{value:function(key){return map[typeof(key)].has(key);},writable:false,configurable:false,enumerable:false},'delete':{value:function(key){if(this.has(key)){size--;keys.splice(keys.indexOf(key),1);return map[typeof(key)]['delete'](key);}
return false;},writable:false,configurable:false,enumerable:false},clear:{value:function(){keys.length=0;for(var key in map){map[key].clear();}
size=0;},writable:false,configurable:false,enumerable:false},forEach:{value:function(callback,thisArg){for(var i=0,n=keys.length;i<n;i++){callback.call(thisArg,this.get(keys[i]),keys[i],this);}},writable:false,configurable:false,enumerable:false}});}
else{m.size=size;m.get=function(key){return map[typeof(key)].get(key);};m.set=function(key,value){if(!this.has(key)){keys.push(key);this.size=++size;}
map[typeof(key)].set(key,value);};m.has=function(key){return map[typeof(key)].has(key);};m['delete']=function(key){if(this.has(key)){this.size=--size;keys.splice(keys.indexOf(key),1);return map[typeof(key)]['delete'](key);}
return false;};m.clear=function(){keys.length=0;for(var key in map){map[key].clear();}
this.size=size=0;};m.forEach=function(callback,thisArg){for(var i=0,n=keys.length;i<n;i++){callback.call(thisArg,this.get(keys[i]),keys[i],this);}};}
if(arguments.length>0&&Array.isArray(arguments[0])){var iterable=arguments[0];for(var i=0,len=iterable.length;i<len;i++){m.set(iterable[i][0],iterable[i][1]);}}
return m;};}
if(!hasForEach){var OldMap=global.Map;global.Map=function Map(){var map=new OldMap();var size=0;var keys=[];var m=Object.create(Map.prototype,{size:{get:function(){return size;},configurable:false,enumerable:false},get:{value:function(key){return map.get(key);},writable:false,configurable:false,enumerable:false},set:{value:function(key,value){if(!map.has(key)){keys.push(key);size++;}
map.set(key,value);},writable:false,configurable:false,enumerable:false},has:{value:function(key){return map.has(key);},writable:false,configurable:false,enumerable:false},'delete':{value:function(key){if(map.has(key)){size--;keys.splice(keys.indexOf(key),1);return map['delete'](key);}
return false;},writable:false,configurable:false,enumerable:false},clear:{value:function(){if('clear'in map){map.clear();}
else{for(var i=0,n=keys.length;i<n;i++){map['delete'](keys[i]);}}
keys.length=0;size=0;},writable:false,configurable:false,enumerable:false},forEach:{value:function(callback,thisArg){for(var i=0,n=keys.length;i<n;i++){callback.call(thisArg,this.get(keys[i]),keys[i],this);}},writable:false,configurable:false,enumerable:false}});if(arguments.length>0&&Array.isArray(arguments[0])){var iterable=arguments[0];for(var i=0,len=iterable.length;i<len;i++){m.set(iterable[i][0],iterable[i][1]);}}
return m;};}})(this||[eval][0]('this'));(function(global){if(typeof global.TimeoutError!=='function'){var TimeoutError=function(message){Error.call(this);this.message=message;this.name=TimeoutError.name;if(typeof Error.captureStackTrace==='function'){Error.captureStackTrace(this,TimeoutError);}}
TimeoutError.prototype=global.hprose.createObject(Error.prototype);TimeoutError.prototype.constructor=TimeoutError;global.TimeoutError=TimeoutError;}})(this||[eval][0]('this'));(function(global,undefined){'use strict';if(global.setImmediate){return;}
var doc=global.document;var MutationObserver=global.MutationObserver||global.WebKitMutationObserver||global.MozMutationOvserver;var polifill={};var nextId=1;var tasks={};function wrap(handler){var args=Array.prototype.slice.call(arguments,1);return function(){handler.apply(undefined,args);};}
function clear(handleId){delete tasks[handleId];}
function run(handleId){var task=tasks[handleId];if(task){try{task();}
finally{clear(handleId);}}}
function create(args){tasks[nextId]=wrap.apply(undefined,args);return nextId++;}
polifill.mutationObserver=function(){var queue=[],node=doc.createTextNode(''),observer=new MutationObserver(function(){while(queue.length>0){run(queue.shift());}});observer.observe(node,{"characterData":true});return function(){var handleId=create(arguments);queue.push(handleId);node.data=handleId&1;return handleId;};};polifill.messageChannel=function(){var channel=new global.MessageChannel();channel.port1.onmessage=function(event){run(Number(event.data));};return function(){var handleId=create(arguments);channel.port2.postMessage(handleId);return handleId;};};polifill.nextTick=function(){return function(){var handleId=create(arguments);global.process.nextTick(wrap(run,handleId));return handleId;};};polifill.postMessage=function(){var iframe=doc.createElement('iframe');iframe.style.display='none';doc.documentElement.appendChild(iframe);var iwin=iframe.contentWindow;iwin.document.write('<script>window.onmessage=function(){parent.postMessage(1, "*");};</script>');iwin.document.close();var queue=[];window.addEventListener('message',function(){while(queue.length>0){run(queue.shift());}});return function(){var handleId=create(arguments);queue.push(handleId);iwin.postMessage(1,"*");return handleId;};};polifill.readyStateChange=function(){var html=doc.documentElement;return function(){var handleId=create(arguments);var script=doc.createElement('script');script.onreadystatechange=function(){run(handleId);script.onreadystatechange=null;html.removeChild(script);script=null;};html.appendChild(script);return handleId;};};var attachTo=Object.getPrototypeOf&&Object.getPrototypeOf(global);attachTo=(attachTo&&attachTo.setTimeout?attachTo:global);polifill.setTimeout=function(){return function(){var handleId=create(arguments);attachTo.setTimeout(wrap(run,handleId),0);return handleId;};};if(typeof(global.process)!=='undefined'&&Object.prototype.toString.call(global.process)==='[object process]'&&!global.process.browser){attachTo.setImmediate=polifill.nextTick();}
else if(doc&&('onreadystatechange'in doc.createElement('script'))){attachTo.setImmediate=polifill.readyStateChange();}
else if(doc&&MutationObserver){attachTo.setImmediate=polifill.mutationObserver();}
else if(global.MessageChannel){attachTo.setImmediate=polifill.messageChannel();}
else if(doc&&'postMessage'in global&&'addEventListener'in global){attachTo.setImmediate=polifill.postMessage();}
else{attachTo.setImmediate=polifill.setTimeout();}
attachTo.clearImmediate=clear;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var PENDING=0;var FULFILLED=1;var REJECTED=2;var defineProperties=global.hprose.defineProperties;var createObject=global.hprose.createObject;var hasPromise='Promise'in global;var setImmediate=global.setImmediate;var setTimeout=global.setTimeout;var clearTimeout=global.clearTimeout;var TimeoutError=global.TimeoutError;function Future(computation){var self=this;defineProperties(self,{_subscribers:{value:[]},resolve:{value:this.resolve.bind(self)},reject:{value:this.reject.bind(self)}});if(typeof computation==='function'){setImmediate(function(){try{self.resolve(computation());}
catch(e){self.reject(e);}});}}
function isFuture(obj){return obj instanceof Future;}
function isPromise(obj){return isFuture(obj)||(hasPromise&&(obj instanceof global.Promise)&&(typeof(obj.then==='function')));}
function toPromise(obj){return(isFuture(obj)?obj:value(obj));}
function delayed(duration,value){var computation=(typeof value==='function')?value:function(){return value;};var future=new Future();setTimeout(function(){try{future.resolve(computation());}
catch(e){future.reject(e);}},duration);return future;}
function error(e){var future=new Future();future.reject(e);return future;}
function value(v){var future=new Future();future.resolve(v);return future;}
function sync(computation){try{var result=computation();return value(result);}
catch(e){return error(e);}}
function promise(executor){var future=new Future();executor(future.resolve,future.reject);return future;}
function arraysize(array){var size=0;Array.forEach(array,function(){++size;});return size;}
function all(array){array=isPromise(array)?array:value(array);return array.then(function(array){var n=array.length;var count=arraysize(array);var result=new Array(n);if(count===0){return value(result);}
var future=new Future();Array.forEach(array,function(element,index){toPromise(element).then(function(value){result[index]=value;if(--count===0){future.resolve(result);}},future.reject);});return future;});}
function join(){return all(arguments);}
function race(array){array=isPromise(array)?array:value(array);return array.then(function(array){var future=new Future();Array.forEach(array,function(element){toPromise(element).fill(future);});return future;});}
function any(array){array=isPromise(array)?array:value(array);return array.then(function(array){var n=array.length;var count=arraysize(array);if(count===0){throw new RangeError('any(): array must not be empty');}
var reasons=new Array(n);var future=new Future();Array.forEach(array,function(element,index){toPromise(element).then(future.resolve,function(e){reasons[index]=e;if(--count===0){future.reject(reasons);}});});return future;});}
function settle(array){array=isPromise(array)?array:value(array);return array.then(function(array){var n=array.length;var count=arraysize(array);var result=new Array(n);if(count===0){return value(result);}
var future=new Future();Array.forEach(array,function(element,index){var f=toPromise(element);f.whenComplete(function(){result[index]=f.inspect();if(--count===0){future.resolve(result);}});});return future;});}
function attempt(handler){var args=Array.slice(arguments,1);return all(args).then(function(args){return handler.apply(undefined,args);});}
function run(handler,thisArg){var args=Array.slice(arguments,2);return all(args).then(function(args){return handler.apply(thisArg,args);});}
function wrap(handler,thisArg){return function(){return all(arguments).then(function(args){return handler.apply(thisArg,args);});};}
function forEach(array,callback,thisArg){return all(array).then(function(array){return array.forEach(callback,thisArg);});}
function every(array,callback,thisArg){return all(array).then(function(array){return array.every(callback,thisArg);});}
function some(array,callback,thisArg){return all(array).then(function(array){return array.some(callback,thisArg);});}
function filter(array,callback,thisArg){return all(array).then(function(array){return array.filter(callback,thisArg);});}
function map(array,callback,thisArg){return all(array).then(function(array){return array.map(callback,thisArg);});}
function reduce(array,callback,initialValue){if(arguments.length>2){return all(array).then(function(array){if(!isPromise(initialValue)){initialValue=value(initialValue);}
return initialValue.then(function(value){return array.reduce(callback,value);});});}
return all(array).then(function(array){return array.reduce(callback);});}
function reduceRight(array,callback,initialValue){if(arguments.length>2){return all(array).then(function(array){if(!isPromise(initialValue)){initialValue=value(initialValue);}
return initialValue.then(function(value){return array.reduceRight(callback,value);});});}
return all(array).then(function(array){return array.reduceRight(callback);});}
function indexOf(array,searchElement,fromIndex){return all(array).then(function(array){if(!isPromise(searchElement)){searchElement=value(searchElement);}
return searchElement.then(function(searchElement){return array.indexOf(searchElement,fromIndex);});});}
function lastIndexOf(array,searchElement,fromIndex){return all(array).then(function(array){if(!isPromise(searchElement)){searchElement=value(searchElement);}
return searchElement.then(function(searchElement){if(fromIndex===undefined){fromIndex=array.length-1;}
return array.lastIndexOf(searchElement,fromIndex);});});}
function includes(array,searchElement,fromIndex){return all(array).then(function(array){if(!isPromise(searchElement)){searchElement=value(searchElement);}
return searchElement.then(function(searchElement){return array.includes(searchElement,fromIndex);});});}
function find(array,predicate,thisArg){return all(array).then(function(array){return array.find(predicate,thisArg);});}
function findIndex(array,predicate,thisArg){return all(array).then(function(array){return array.findIndex(predicate,thisArg);});}
defineProperties(Future,{delayed:{value:delayed},error:{value:error},sync:{value:sync},value:{value:value},all:{value:all},race:{value:race},resolve:{value:value},reject:{value:error},promise:{value:promise},isFuture:{value:isFuture},isPromise:{value:isPromise},toPromise:{value:toPromise},join:{value:join},any:{value:any},settle:{value:settle},attempt:{value:attempt},run:{value:run},wrap:{value:wrap},forEach:{value:forEach},every:{value:every},some:{value:some},filter:{value:filter},map:{value:map},reduce:{value:reduce},reduceRight:{value:reduceRight},indexOf:{value:indexOf},lastIndexOf:{value:lastIndexOf},includes:{value:includes},find:{value:find},findIndex:{value:findIndex}});function _call(callback,next,x){setImmediate(function(){try{var r=callback(x);next.resolve(r);}
catch(e){next.reject(e);}});}
function _resolve(onfulfill,next,x){if(onfulfill){_call(onfulfill,next,x);}
else{next.resolve(x);}}
function _reject(onreject,next,e){if(onreject){_call(onreject,next,e);}
else{next.reject(e);}}
defineProperties(Future.prototype,{_value:{writable:true},_reason:{writable:true},_state:{value:PENDING,writable:true},resolve:{value:function(value){if(value===this){this.reject(new TypeError('Self resolution'));return;}
if(isFuture(value)){value.fill(this);return;}
if((value!==null)&&(typeof value==='object')||(typeof value==='function')){var then;try{then=value.then;}
catch(e){this.reject(e);return;}
if(typeof then==='function'){var notrun=true;try{var self=this;then.call(value,function(y){if(notrun){notrun=false;self.resolve(y);}},function(r){if(notrun){notrun=false;self.reject(r);}});return;}
catch(e){if(notrun){notrun=false;this.reject(e);}}
return;}}
if(this._state===PENDING){this._state=FULFILLED;this._value=value;var subscribers=this._subscribers;while(subscribers.length>0){var subscriber=subscribers.shift();_resolve(subscriber.onfulfill,subscriber.next,value);}}}},reject:{value:function(reason){if(this._state===PENDING){this._state=REJECTED;this._reason=reason;var subscribers=this._subscribers;while(subscribers.length>0){var subscriber=subscribers.shift();_reject(subscriber.onreject,subscriber.next,reason);}}}},then:{value:function(onfulfill,onreject){if(typeof onfulfill!=='function'){onfulfill=null;}
if(typeof onreject!=='function'){onreject=null;}
var next=new Future();if(this._state===FULFILLED){_resolve(onfulfill,next,this._value);}
else if(this._state===REJECTED){_reject(onreject,next,this._reason);}
else{this._subscribers.push({onfulfill:onfulfill,onreject:onreject,next:next});}
return next;}},done:{value:function(onfulfill,onreject){this.then(onfulfill,onreject).then(null,function(error){setImmediate(function(){throw error;});});}},inspect:{value:function(){switch(this._state){case PENDING:return{state:'pending'};case FULFILLED:return{state:'fulfilled',value:this._value};case REJECTED:return{state:'rejected',reason:this._reason};}}},catchError:{value:function(onreject,test){if(typeof test==='function'){var self=this;return this['catch'](function(e){if(test(e)){return self['catch'](onreject);}
else{throw e;}});}
return this['catch'](onreject);}},'catch':{value:function(onreject){return this.then(null,onreject);}},fail:{value:function(onreject){this.done(null,onreject);}},whenComplete:{value:function(action){return this.then(function(v){action();return v;},function(e){action();throw e;});}},complete:{value:function(oncomplete){return this.then(oncomplete,oncomplete);}},always:{value:function(oncomplete){this.done(oncomplete,oncomplete);}},fill:{value:function(future){this.then(future.resolve,future.reject);}},timeout:{value:function(duration,reason){var future=new Future();var timeoutId=setTimeout(function(){future.reject(reason||new TimeoutError('timeout'));},duration);this.whenComplete(function(){clearTimeout(timeoutId);}).fill(future);return future;}},delay:{value:function(duration){var future=new Future();this.then(function(result){setTimeout(function(){future.resolve(result);},duration);},future.reject);return future;}},tap:{value:function(onfulfilledSideEffect,thisArg){return this.then(function(result){onfulfilledSideEffect.call(thisArg,result);return result;});}},spread:{value:function(onfulfilledArray,thisArg){return this.then(function(array){return onfulfilledArray.apply(thisArg,array);});}},get:{value:function(key){return this.then(function(result){return result[key];});}},set:{value:function(key,value){return this.then(function(result){result[key]=value;return result;});}},apply:{value:function(method,args){args=args||[];return this.then(function(result){return all(args).then(function(args){return result[method].apply(result,args);});});}},call:{value:function(method){var args=Array.slice(arguments,1);return this.then(function(result){return all(args).then(function(args){return result[method].apply(result,args);});});}},bind:{value:function(method){var bindargs=Array.slice(arguments);if(Array.isArray(method)){for(var i=0,n=method.length;i<n;++i){bindargs[0]=method[i];this.bind.apply(this,bindargs);}
return;}
bindargs.shift();var self=this;Object.defineProperty(this,method,{value:function(){var args=Array.slice(arguments);return self.then(function(result){return all(bindargs.concat(args)).then(function(args){return result[method].apply(result,args);});});}});return this;}},forEach:{value:function(callback,thisArg){return forEach(this,callback,thisArg);}},every:{value:function(callback,thisArg){return every(this,callback,thisArg);}},some:{value:function(callback,thisArg){return some(this,callback,thisArg);}},filter:{value:function(callback,thisArg){return filter(this,callback,thisArg);}},map:{value:function(callback,thisArg){return map(this,callback,thisArg);}},reduce:{value:function(callback,initialValue){if(arguments.length>1){return reduce(this,callback,initialValue);}
return reduce(this,callback);}},reduceRight:{value:function(callback,initialValue){if(arguments.length>1){return reduceRight(this,callback,initialValue);}
return reduceRight(this,callback);}},indexOf:{value:function(searchElement,fromIndex){return indexOf(this,searchElement,fromIndex);}},lastIndexOf:{value:function(searchElement,fromIndex){return lastIndexOf(this,searchElement,fromIndex);}},includes:{value:function(searchElement,fromIndex){return includes(this,searchElement,fromIndex);}},find:{value:function(predicate,thisArg){return find(this,predicate,thisArg);}},findIndex:{value:function(predicate,thisArg){return findIndex(this,predicate,thisArg);}}});global.hprose.Future=Future;function Completer(){var future=new Future();defineProperties(this,{future:{value:future},complete:{value:future.resolve},completeError:{value:future.reject},isCompleted:{get:function(){return(future._state!==PENDING);}}});}
global.hprose.Completer=Completer;global.hprose.resolved=value;global.hprose.rejected=error;global.hprose.deferred=function(){var self=new Future();return createObject(null,{promise:{value:self},resolve:{value:self.resolve},reject:{value:self.reject}});};if(hasPromise){return;}
global.Promise=function(executor){Future.call(this);executor(this.resolve,this.reject);};global.Promise.prototype=createObject(Future.prototype);global.Promise.prototype.constructor=Future;defineProperties(global.Promise,{all:{value:all},race:{value:race},resolve:{value:value},reject:{value:error}});})(this||[eval][0]('this'));(function(global){'use strict';var defineProperties=global.hprose.defineProperties;var createObject=global.hprose.createObject;function BinaryString(bs,needtest){if(!needtest||/^[\x00-\xff]*$/.test(bs)){defineProperties(this,{length:{value:bs.length},toString:{value:function(){return bs;}},valueOf:{value:function(){return bs;},writable:true,configurable:true,enumerable:false}});}
else{throw new Error("argument is not a binary string.");}}
var methods={};['quote','substring','toLowerCase','toUpperCase','charAt','charCodeAt','indexOf','lastIndexOf','include','startsWith','endsWith','repeat','trim','trimLeft','trimRight','toLocaleLowerCase','toLocaleUpperCase','match','search','replace','split','substr','concat','slice'].forEach(function(name){methods[name]={value:String.prototype[name]};});BinaryString.prototype=createObject(null,methods);BinaryString.prototype.constructor=BinaryString;global.hprose.BinaryString=BinaryString;global.hprose.binary=function(bs){return new BinaryString(bs,true);};})(this||[eval][0]('this'));(function(global,undefined){'use strict';var defineProperties=global.hprose.defineProperties;function int32BE(i){return String.fromCharCode(i>>>24&0xFF,i>>>16&0xFF,i>>>8&0xFF,i&0xFF);}
function int32LE(i){return String.fromCharCode(i&0xFF,i>>>8&0xFF,i>>>16&0xFF,i>>>24&0xFF);}
function utf8Encode(s){var buf=[];var n=s.length;for(var i=0,j=0;i<n;++i,++j){var codeUnit=s.charCodeAt(i);if(codeUnit<0x80){buf[j]=s.charAt(i);}
else if(codeUnit<0x800){buf[j]=String.fromCharCode(0xC0|(codeUnit>>6),0x80|(codeUnit&0x3F));}
else if(codeUnit<0xD800||codeUnit>0xDFFF){buf[j]=String.fromCharCode(0xE0|(codeUnit>>12),0x80|((codeUnit>>6)&0x3F),0x80|(codeUnit&0x3F));}
else{if(i+1<n){var nextCodeUnit=s.charCodeAt(i+1);if(codeUnit<0xDC00&&0xDC00<=nextCodeUnit&&nextCodeUnit<=0xDFFF){var rune=(((codeUnit&0x03FF)<<10)|(nextCodeUnit&0x03FF))+0x010000;buf[j]=String.fromCharCode(0xF0|((rune>>18)&0x3F),0x80|((rune>>12)&0x3F),0x80|((rune>>6)&0x3F),0x80|(rune&0x3F));++i;continue;}}
throw new Error('Malformed string');}}
return buf.join('');}
function readShortString(bs,n){var charCodes=new Array(n);var i=0,off=0;for(var len=bs.length;i<n&&off<len;i++){var unit=bs.charCodeAt(off++);switch(unit>>4){case 0:case 1:case 2:case 3:case 4:case 5:case 6:case 7:charCodes[i]=unit;break;case 12:case 13:if(off<len){charCodes[i]=((unit&0x1F)<<6)|(bs.charCodeAt(off++)&0x3F);}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 14:if(off+1<len){charCodes[i]=((unit&0x0F)<<12)|((bs.charCodeAt(off++)&0x3F)<<6)|(bs.charCodeAt(off++)&0x3F);}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 15:if(off+2<len){var rune=(((unit&0x07)<<18)|((bs.charCodeAt(off++)&0x3F)<<12)|((bs.charCodeAt(off++)&0x3F)<<6)|(bs.charCodeAt(off++)&0x3F))-0x10000;if(0<=rune&&rune<=0xFFFFF){charCodes[i++]=(((rune>>10)&0x03FF)|0xD800);charCodes[i]=((rune&0x03FF)|0xDC00);}
else{throw new Error('Character outside valid Unicode range: 0x'+rune.toString(16));}}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;default:throw new Error('Bad UTF-8 encoding 0x'+unit.toString(16));}}
if(i<n){charCodes.length=i;}
return[String.fromCharCode.apply(String,charCodes),off];}
function readLongString(bs,n){var buf=[];var charCodes=new Array(0x8000);var i=0,off=0;for(var len=bs.length;i<n&&off<len;i++){var unit=bs.charCodeAt(off++);switch(unit>>4){case 0:case 1:case 2:case 3:case 4:case 5:case 6:case 7:charCodes[i]=unit;break;case 12:case 13:if(off<len){charCodes[i]=((unit&0x1F)<<6)|(bs.charCodeAt(off++)&0x3F);}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 14:if(off+1<len){charCodes[i]=((unit&0x0F)<<12)|((bs.charCodeAt(off++)&0x3F)<<6)|(bs.charCodeAt(off++)&0x3F);}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 15:if(off+2<len){var rune=(((unit&0x07)<<18)|((bs.charCodeAt(off++)&0x3F)<<12)|((bs.charCodeAt(off++)&0x3F)<<6)|(bs.charCodeAt(off++)&0x3F))-0x10000;if(0<=rune&&rune<=0xFFFFF){charCodes[i++]=(((rune>>10)&0x03FF)|0xD800);charCodes[i]=((rune&0x03FF)|0xDC00);}
else{throw new Error('Character outside valid Unicode range: 0x'+rune.toString(16));}}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;default:throw new Error('Bad UTF-8 encoding 0x'+unit.toString(16));}
if(i>=0x7FFF-1){var size=i+1;charCodes.length=size;buf[buf.length]=String.fromCharCode.apply(String,charCodes);n-=size;i=-1;}}
if(i>0){charCodes.length=i;buf[buf.length]=String.fromCharCode.apply(String,charCodes);}
return[buf.join(''),off];}
function readString(bs,n){if(n===undefined||n===null||(n<0)){n=bs.length;}
if(n===0){return['',0];}
return((n<0xFFFF)?readShortString(bs,n):readLongString(bs,n));}
function readUTF8(bs,n){if(n===undefined||n===null||(n<0)){n=bs.length;}
if(n===0){return'';}
var i=0,off=0;for(var len=bs.length;i<n&&off<len;i++){var unit=bs.charCodeAt(off++);switch(unit>>4){case 0:case 1:case 2:case 3:case 4:case 5:case 6:case 7:break;case 12:case 13:if(off<len){++off;}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 14:if(off+1<len){off+=2;}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 15:if(off+2<len){var rune=(((unit&0x07)<<18)|((bs.charCodeAt(off++)&0x3F)<<12)|((bs.charCodeAt(off++)&0x3F)<<6)|(bs.charCodeAt(off++)&0x3F))-0x10000;if(0<=rune&&rune<=0xFFFFF){break;}
throw new Error('Character outside valid Unicode range: 0x'+rune.toString(16));}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;default:throw new Error('Bad UTF-8 encoding 0x'+unit.toString(16));}}
return bs.substr(0,off);}
function utf8Decode(bs){return readString(bs)[0];}
function utf8Length(s){var n=s.length;var length=0;for(var i=0;i<n;++i){var codeUnit=s.charCodeAt(i);if(codeUnit<0x80){++length;}
else if(codeUnit<0x800){length+=2;}
else if(codeUnit<0xD800||codeUnit>0xDFFF){length+=3;}
else{if(i+1<n){var nextCodeUnit=s.charCodeAt(i+1);if(codeUnit<0xDC00&&0xDC00<=nextCodeUnit&&nextCodeUnit<=0xDFFF){++i;length+=4;continue;}}
throw new Error('Malformed string');}}
return length;}
function utf16Length(bs){var n=bs.length;var length=0;for(var i=0;i<n;++i,++length){var unit=bs.charCodeAt(i);switch(unit>>4){case 0:case 1:case 2:case 3:case 4:case 5:case 6:case 7:break;case 12:case 13:if(i<n){++i;}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 14:if(i+1<n){i+=2;}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;case 15:if(i+2<n){var rune=(((unit&0x07)<<18)|((bs.charCodeAt(i++)&0x3F)<<12)|((bs.charCodeAt(i++)&0x3F)<<6)|(bs.charCodeAt(i++)&0x3F))-0x10000;if(0<=rune&&rune<=0xFFFFF){++length;}
else{throw new Error('Character outside valid Unicode range: 0x'+rune.toString(16));}}
else{throw new Error('Unfinished UTF-8 octet sequence');}
break;default:throw new Error('Bad UTF-8 encoding 0x'+unit.toString(16));}}
return length;}
function isUTF8(bs){for(var i=0,n=bs.length;i<n;++i){var unit=bs.charCodeAt(i);switch(unit>>4){case 0:case 1:case 2:case 3:case 4:case 5:case 6:case 7:break;case 12:case 13:if(i<n){++i;}
else{return false;}
break;case 14:if(i+1<n){i+=2;}
else{return false;}
break;case 15:if(i+2<n){var rune=(((unit&0x07)<<18)|((bs.charCodeAt(i++)&0x3F)<<12)|((bs.charCodeAt(i++)&0x3F)<<6)|(bs.charCodeAt(i++)&0x3F))-0x10000;if(!(0<=rune&&rune<=0xFFFFF)){return false;}}
else{return false;}
break;default:return false;}}
return true;}
function StringIO(){var a=arguments;switch(a.length){case 1:this._buffer=[a[0].toString()];break;case 2:this._buffer=[a[0].toString().substr(a[1])];break;case 3:this._buffer=[a[0].toString().substr(a[1],a[2])];break;default:this._buffer=[''];break;}
this.mark();}
defineProperties(StringIO.prototype,{_buffer:{writable:true},_off:{value:0,writable:true},_wmark:{writable:true},_rmark:{writable:true},toString:{value:function(){if(this._buffer.length>1){this._buffer=[this._buffer.join('')];}
return this._buffer[0];}},length:{get:function(){return this.toString().length;}},position:{get:function(){return this._off;}},mark:{value:function(){this._wmark=this.length();this._rmark=this._off;}},reset:{value:function(){this._buffer=[this.toString().substr(0,this._wmark)];this._off=this._rmark;}},clear:{value:function(){this._buffer=[''];this._wmark=0;this._off=0;this._rmark=0;}},writeByte:{value:function(b){this._buffer.push(String.fromCharCode(b&0xFF));}},writeInt32BE:{value:function(i){if((i===(i|0))&&(i<=2147483647)){this._buffer.push(int32BE(i));return;}
throw new TypeError('value is out of bounds');}},writeUInt32BE:{value:function(i){if(((i&0x7FFFFFFF)+0x80000000===i)&&(i>=0)){this._buffer.push(int32BE(i|0));return;}
throw new TypeError('value is out of bounds');}},writeInt32LE:{value:function(i){if((i===(i|0))&&(i<=2147483647)){this._buffer.push(int32LE(i));return;}
throw new TypeError('value is out of bounds');}},writeUInt32LE:{value:function(i){if(((i&0x7FFFFFFF)+0x80000000===i)&&(i>=0)){this._buffer.push(int32LE(i|0));return;}
throw new TypeError('value is out of bounds');}},writeUTF16AsUTF8:{value:function(str){this._buffer.push(utf8Encode(str));}},writeUTF8AsUTF16:{value:function(str){this._buffer.push(utf8Decode(str));}},write:{value:function(data){this._buffer.push(data);}},readByte:{value:function(){if(this._off<this.length()){return this._buffer[0].charCodeAt(this._off++);}
return-1;}},readChar:{value:function(){if(this._off<this.length()){return this._buffer[0].charAt(this._off++);}
return'';}},readInt32BE:{value:function(){var len=this.length();var buf=this._buffer[0];var off=this._off;if(off+3<len){var result=buf.charCodeAt(off++)<<24|buf.charCodeAt(off++)<<16|buf.charCodeAt(off++)<<8|buf.charCodeAt(off++);this._off=off;return result;}
throw new Error('EOF');}},readUInt32BE:{value:function(){var value=this.readInt32BE();if(value<0){return(value&0x7FFFFFFF)+0x80000000;}
return value;}},readInt32LE:{value:function(){var len=this.length();var buf=this._buffer[0];var off=this._off;if(off+3<len){var result=buf.charCodeAt(off++)|buf.charCodeAt(off++)<<8|buf.charCodeAt(off++)<<16|buf.charCodeAt(off++)<<24;this._off=off;return result;}
throw new Error('EOF');}},readUInt32LE:{value:function(){var value=this.readInt32LE();if(value<0){return(value&0x7FFFFFFF)+0x80000000;}
return value;}},read:{value:function(n){var off=this._off;var len=this.length();if(off+n>len){n=len-off;}
if(n===0){return'';}
this._off=off+n;return this._buffer[0].substring(off,this._off);}},skip:{value:function(n){var len=this.length();if(this._off+n>len){n=len-this._off;this._off=len;}
else{this._off+=n;}
return n;}},readString:{value:function(tag){var len=this.length();var off=this._off;var buf=this._buffer[0];var pos=buf.indexOf(tag,off);if(pos===-1){buf=buf.substr(off);this._off=len;}
else{buf=buf.substring(off,pos+1);this._off=pos+1;}
return buf;}},readUntil:{value:function(tag){var len=this.length();var off=this._off;var buf=this._buffer[0];var pos=buf.indexOf(tag,off);if(pos===this._off){buf='';this._off++;}
else if(pos===-1){buf=buf.substr(off);this._off=len;}
else{buf=buf.substring(off,pos);this._off=pos+1;}
return buf;}},readUTF8:{value:function(n){var len=this.length();var r=readUTF8(this._buffer[0].substring(this._off,Math.min(this._off+n*3,len)),n);this._off+=r.length;return r;}},readUTF8AsUTF16:{value:function(n){var len=this.length();var r=readString(this._buffer[0].substring(this._off,Math.min(this._off+n*3,len)),n);this._off+=r[1];return r[0];}},readUTF16AsUTF8:{value:function(n){return utf8Encode(this.read(n));}},take:{value:function(){var buffer=this.toString();this.clear();return buffer;}},clone:{value:function(){return new StringIO(this.toString());}},trunc:{value:function(){var buf=this.toString().substring(this._off,this._length);this._buffer[0]=buf;this._off=0;this._wmark=0;this._rmark=0;}}});defineProperties(StringIO,{utf8Encode:{value:utf8Encode},utf8Decode:{value:utf8Decode},utf8Length:{value:utf8Length},utf16Length:{value:utf16Length},isUTF8:{value:isUTF8}});global.hprose.StringIO=StringIO;})(this||[eval][0]('this'));(function(global){'use strict';global.HproseTags=global.hprose.Tags={TagInteger:'i',TagLong:'l',TagDouble:'d',TagNull:'n',TagEmpty:'e',TagTrue:'t',TagFalse:'f',TagNaN:'N',TagInfinity:'I',TagDate:'D',TagTime:'T',TagUTC:'Z',TagBytes:'b',TagUTF8Char:'u',TagString:'s',TagGuid:'g',TagList:'a',TagMap:'m',TagClass:'c',TagObject:'o',TagRef:'r',TagPos:'+',TagNeg:'-',TagSemicolon:';',TagOpenbrace:'{',TagClosebrace:'}',TagQuote:'"',TagPoint:'.',TagFunctions:'F',TagCall:'C',TagResult:'R',TagArgument:'A',TagError:'E',TagEnd:'z'};})(this||[eval][0]('this'));(function(global){'use strict';var WeakMap=global.WeakMap;var createObject=global.hprose.createObject;var classCache=createObject(null);var aliasCache=new WeakMap();function register(cls,alias){aliasCache.set(cls,alias);classCache[alias]=cls;}
function getClassAlias(cls){return aliasCache.get(cls);}
function getClass(alias){return classCache[alias];}
global.HproseClassManager=global.hprose.ClassManager=createObject(null,{register:{value:register},getClassAlias:{value:getClassAlias},getClass:{value:getClass}});global.hprose.register=register;register(Object,'Object');})(this||[eval][0]('this'));(function(global,undefined){'use strict';var Map=global.Map;var StringIO=global.hprose.StringIO;var BinaryString=global.hprose.BinaryString;var Tags=global.hprose.Tags;var ClassManager=global.hprose.ClassManager;var defineProperties=global.hprose.defineProperties;var createObject=global.hprose.createObject;var utf8Encode=StringIO.utf8Encode;function getClassName(obj){var cls=obj.constructor;var classname=ClassManager.getClassAlias(cls);if(classname){return classname;}
if(cls.name){classname=cls.name;}
else{var ctor=cls.toString();classname=ctor.substr(0,ctor.indexOf('(')).replace(/(^\s*function\s*)|(\s*$)/ig,'');if(classname===''||classname==='Object'){return(typeof(obj.getClassName)==='function')?obj.getClassName():'Object';}}
if(classname!=='Object'){ClassManager.register(cls,classname);}
return classname;}
var fakeWriterRefer=createObject(null,{set:{value:function(){}},write:{value:function(){return false;}},reset:{value:function(){}}});function RealWriterRefer(stream){defineProperties(this,{_stream:{value:stream},_ref:{value:new Map(),writable:true}});}
defineProperties(RealWriterRefer.prototype,{_refcount:{value:0,writable:true},set:{value:function(val){this._ref.set(val,this._refcount++);}},write:{value:function(val){var index=this._ref.get(val);if(index!==undefined){this._stream.write(Tags.TagRef);this._stream.write(index);this._stream.write(Tags.TagSemicolon);return true;}
return false;}},reset:{value:function(){this._ref=new Map();this._refcount=0;}}});function realWriterRefer(stream){return new RealWriterRefer(stream);}
function Writer(stream,simple,binary){this.binary=!!binary;defineProperties(this,{stream:{value:stream},_classref:{value:createObject(null),writable:true},_fieldsref:{value:[],writable:true},_refer:{value:simple?fakeWriterRefer:realWriterRefer(stream)}});}
function serialize(writer,value){var stream=writer.stream;if(value===undefined||value===null||value.constructor===Function){stream.write(Tags.TagNull);return;}
if(value===''){stream.write(Tags.TagEmpty);return;}
switch(value.constructor){case Number:writeNumber(writer,value);break;case Boolean:writeBoolean(writer,value);break;case String:if(value.length===1){stream.write(Tags.TagUTF8Char);stream.write(writer.binary?utf8Encode(value):value);}
else{writer.writeStringWithRef(value);}
break;case BinaryString:if(writer.binary){writer.writeBinaryWithRef(value);}
else{throw new Error('The binary string does not support serialization in text mode.');}
break;case Date:writer.writeDateWithRef(value);break;case Map:writer.writeMapWithRef(value);break;default:if(Array.isArray(value)){writer.writeListWithRef(value);}
else{var classname=getClassName(value);if(classname==='Object'){writer.writeMapWithRef(value);}
else{writer.writeObjectWithRef(value);}}
break;}}
function writeNumber(writer,n){var stream=writer.stream;n=n.valueOf();if(n===(n|0)){if(0<=n&&n<=9){stream.write(n);}
else{stream.write(Tags.TagInteger);stream.write(n);stream.write(Tags.TagSemicolon);}}
else{writeDouble(writer,n);}}
function writeInteger(writer,n){var stream=writer.stream;if(0<=n&&n<=9){stream.write(n);}
else{if(n<-2147483648||n>2147483647){stream.write(Tags.TagLong);}
else{stream.write(Tags.TagInteger);}
stream.write(n);stream.write(Tags.TagSemicolon);}}
function writeDouble(writer,n){var stream=writer.stream;if(n!==n){stream.write(Tags.TagNaN);}
else if(n!==Infinity&&n!==-Infinity){stream.write(Tags.TagDouble);stream.write(n);stream.write(Tags.TagSemicolon);}
else{stream.write(Tags.TagInfinity);stream.write((n>0)?Tags.TagPos:Tags.TagNeg);}}
function writeBoolean(writer,b){writer.stream.write(b.valueOf()?Tags.TagTrue:Tags.TagFalse);}
function writeUTCDate(writer,date){writer._refer.set(date);var stream=writer.stream;stream.write(Tags.TagDate);stream.write(('0000'+date.getUTCFullYear()).slice(-4));stream.write(('00'+(date.getUTCMonth()+1)).slice(-2));stream.write(('00'+date.getUTCDate()).slice(-2));stream.write(Tags.TagTime);stream.write(('00'+date.getUTCHours()).slice(-2));stream.write(('00'+date.getUTCMinutes()).slice(-2));stream.write(('00'+date.getUTCSeconds()).slice(-2));var millisecond=date.getUTCMilliseconds();if(millisecond!==0){stream.write(Tags.TagPoint);stream.write(('000'+millisecond).slice(-3));}
stream.write(Tags.TagUTC);}
function writeDate(writer,date){writer._refer.set(date);var stream=writer.stream;var year=('0000'+date.getFullYear()).slice(-4);var month=('00'+(date.getMonth()+1)).slice(-2);var day=('00'+date.getDate()).slice(-2);var hour=('00'+date.getHours()).slice(-2);var minute=('00'+date.getMinutes()).slice(-2);var second=('00'+date.getSeconds()).slice(-2);var millisecond=('000'+date.getMilliseconds()).slice(-3);if((hour==='00')&&(minute==='00')&&(second==='00')&&(millisecond==='000')){stream.write(Tags.TagDate);stream.write(year);stream.write(month);stream.write(day);}
else if((year==='1970')&&(month==='01')&&(day==='01')){stream.write(Tags.TagTime);stream.write(hour);stream.write(minute);stream.write(second);if(millisecond!=='000'){stream.write(Tags.TagPoint);stream.write(millisecond);}}
else{stream.write(Tags.TagDate);stream.write(year);stream.write(month);stream.write(day);stream.write(Tags.TagTime);stream.write(hour);stream.write(minute);stream.write(second);if(millisecond!=='000'){stream.write(Tags.TagPoint);stream.write(millisecond);}}
stream.write(Tags.TagSemicolon);}
function writeTime(writer,time){writer._refer.set(time);var stream=writer.stream;var hour=('00'+time.getHours()).slice(-2);var minute=('00'+time.getMinutes()).slice(-2);var second=('00'+time.getSeconds()).slice(-2);var millisecond=('000'+time.getMilliseconds()).slice(-3);stream.write(Tags.TagTime);stream.write(hour);stream.write(minute);stream.write(second);if(millisecond!=='000'){stream.write(Tags.TagPoint);stream.write(millisecond);}
stream.write(Tags.TagSemicolon);}
function writeBinary(writer,bs){writer._refer.set(bs);var stream=writer.stream;stream.write(Tags.TagBytes);var n=bs.length;if(n>0){stream.write(n);stream.write(Tags.TagQuote);stream.write(bs);}
else{stream.write(Tags.TagQuote);}
stream.write(Tags.TagQuote);}
function writeString(writer,str){writer._refer.set(str);var stream=writer.stream;var n=str.length;stream.write(Tags.TagString);if(n>0){stream.write(n);stream.write(Tags.TagQuote);stream.write(writer.binary?utf8Encode(str):str);}
else{stream.write(Tags.TagQuote);}
stream.write(Tags.TagQuote);}
function writeList(writer,array){writer._refer.set(array);var stream=writer.stream;var n=array.length;stream.write(Tags.TagList);if(n>0){stream.write(n);stream.write(Tags.TagOpenbrace);for(var i=0;i<n;i++){serialize(writer,array[i]);}}
else{stream.write(Tags.TagOpenbrace);}
stream.write(Tags.TagClosebrace);}
function writeMap(writer,map){writer._refer.set(map);var stream=writer.stream;var fields=[];for(var key in map){if(map.hasOwnProperty(key)&&typeof(map[key])!=='function'){fields[fields.length]=key;}}
var n=fields.length;stream.write(Tags.TagMap);if(n>0){stream.write(n);stream.write(Tags.TagOpenbrace);for(var i=0;i<n;i++){serialize(writer,fields[i]);serialize(writer,map[fields[i]]);}}
else{stream.write(Tags.TagOpenbrace);}
stream.write(Tags.TagClosebrace);}
function writeHarmonyMap(writer,map){writer._refer.set(map);var stream=writer.stream;var n=map.size;stream.write(Tags.TagMap);if(n>0){stream.write(n);stream.write(Tags.TagOpenbrace);map.forEach(function(value,key){serialize(writer,key);serialize(writer,value);});}
else{stream.write(Tags.TagOpenbrace);}
stream.write(Tags.TagClosebrace);}
function writeObject(writer,obj){var stream=writer.stream;var classname=getClassName(obj);var fields,index;if(classname in writer._classref){index=writer._classref[classname];fields=writer._fieldsref[index];}
else{fields=[];for(var key in obj){if(obj.hasOwnProperty(key)&&typeof(obj[key])!=='function'){fields[fields.length]=key.toString();}}
index=writeClass(writer,classname,fields);}
stream.write(Tags.TagObject);stream.write(index);stream.write(Tags.TagOpenbrace);writer._refer.set(obj);var n=fields.length;for(var i=0;i<n;i++){serialize(writer,obj[fields[i]]);}
stream.write(Tags.TagClosebrace);}
function writeClass(writer,classname,fields){var stream=writer.stream;var n=fields.length;stream.write(Tags.TagClass);stream.write(classname.length);stream.write(Tags.TagQuote);stream.write(writer.binary?utf8Encode(classname):classname);stream.write(Tags.TagQuote);if(n>0){stream.write(n);stream.write(Tags.TagOpenbrace);for(var i=0;i<n;i++){writeString(writer,fields[i]);}}
else{stream.write(Tags.TagOpenbrace);}
stream.write(Tags.TagClosebrace);var index=writer._fieldsref.length;writer._classref[classname]=index;writer._fieldsref[index]=fields;return index;}
defineProperties(Writer.prototype,{binary:{value:false,writable:true},serialize:{value:function(value){serialize(this,value);}},writeInteger:{value:function(value){writeInteger(this,value);}},writeDouble:{value:function(value){writeDouble(this,value);}},writeBoolean:{value:function(value){writeBoolean(this,value);}},writeUTCDate:{value:function(value){writeUTCDate(this,value);}},writeUTCDateWithRef:{value:function(value){if(!this._refer.write(value)){writeUTCDate(this,value);}}},writeDate:{value:function(value){writeDate(this,value);}},writeDateWithRef:{value:function(value){if(!this._refer.write(value)){writeDate(this,value);}}},writeTime:{value:function(value){writeTime(this,value);}},writeTimeWithRef:{value:function(value){if(!this._refer.write(value)){writeTime(this,value);}}},writeBinary:{value:function(value){writeBinary(this,value);}},writeBinaryWithRef:{value:function(value){if(!this._refer.write(value)){writeBinary(this,value);}}},writeString:{value:function(value){writeString(this,value);}},writeStringWithRef:{value:function(value){if(!this._refer.write(value)){writeString(this,value);}}},writeList:{value:function(value){writeList(this,value);}},writeListWithRef:{value:function(value){if(!this._refer.write(value)){writeList(this,value);}}},writeMap:{value:function(value){if(value instanceof Map){writeHarmonyMap(this,value);}
else{writeMap(this,value);}}},writeMapWithRef:{value:function(value){if(!this._refer.write(value)){this.writeMap(value);}}},writeObject:{value:function(value){writeObject(this,value);}},writeObjectWithRef:{value:function(value){if(!this._refer.write(value)){writeObject(this,value);}}},reset:{value:function(){this._classref=createObject(null);this._fieldsref.length=0;this._refer.reset();}}});global.HproseWriter=global.hprose.Writer=Writer;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var Map=global.Map;var StringIO=global.hprose.StringIO;var BinaryString=global.hprose.BinaryString;var Tags=global.hprose.Tags;var ClassManager=global.hprose.ClassManager;var defineProperties=global.hprose.defineProperties;var createObject=global.hprose.createObject;function unexpectedTag(tag,expectTags){if(tag&&expectTags){throw new Error('Tag "'+expectTags+'" expected, but "'+tag+'" found in stream');}
if(tag){throw new Error('Unexpected serialize tag "'+tag+'" in stream');}
throw new Error('No byte found in stream');}
function readRaw(stream,binary){var ostream=new StringIO();_readRaw(stream,ostream,binary);return ostream.take();}
function _readRaw(stream,ostream,binary){__readRaw(stream,ostream,stream.readChar(),binary);}
function __readRaw(stream,ostream,tag,binary){ostream.write(tag);switch(tag){case'0':case'1':case'2':case'3':case'4':case'5':case'6':case'7':case'8':case'9':case Tags.TagNull:case Tags.TagEmpty:case Tags.TagTrue:case Tags.TagFalse:case Tags.TagNaN:break;case Tags.TagInfinity:ostream.write(stream.read());break;case Tags.TagInteger:case Tags.TagLong:case Tags.TagDouble:case Tags.TagRef:readNumberRaw(stream,ostream);break;case Tags.TagDate:case Tags.TagTime:readDateTimeRaw(stream,ostream);break;case Tags.TagUTF8Char:readUTF8CharRaw(stream,ostream,binary);break;case Tags.TagBytes:readBinaryRaw(stream,ostream,binary);break;case Tags.TagString:readStringRaw(stream,ostream,binary);break;case Tags.TagGuid:readGuidRaw(stream,ostream);break;case Tags.TagList:case Tags.TagMap:case Tags.TagObject:readComplexRaw(stream,ostream,binary);break;case Tags.TagClass:readComplexRaw(stream,ostream,binary);_readRaw(stream,ostream,binary);break;case Tags.TagError:_readRaw(stream,ostream,binary);break;default:unexpectedTag(tag);}}
function readNumberRaw(stream,ostream){var tag;do{tag=stream.read();ostream.write(tag);}while(tag!==Tags.TagSemicolon);}
function readDateTimeRaw(stream,ostream){var tag;do{tag=stream.read();ostream.write(tag);}while(tag!==Tags.TagSemicolon&&tag!==Tags.TagUTC);}
function readUTF8CharRaw(stream,ostream,binary){if(binary){ostream.write(stream.readUTF8(1));}
else{ostream.write(stream.readChar());}}
function readBinaryRaw(stream,ostream,binary){if(!binary){throw new Error('The binary string does not support to unserialize in text mode.');}
var s=stream.readUntil(Tags.TagQuote);ostream.write(s);ostream.write(Tags.TagQuote);var count=0;if(s.length>0){count=parseInt(s,10);}
ostream.write(stream.read(count+1));}
function readStringRaw(stream,ostream,binary){var s=stream.readUntil(Tags.TagQuote);ostream.write(s);ostream.write(Tags.TagQuote);var count=0;if(s.length>0){count=parseInt(s,10);}
if(binary){ostream.write(stream.readUTF8(count+1));}
else{ostream.write(stream.read(count+1));}}
function readGuidRaw(stream,ostream){ostream.write(stream.read(38));}
function readComplexRaw(stream,ostream,binary){var tag;do{tag=stream.readChar();ostream.write(tag);}while(tag!==Tags.TagOpenbrace);while((tag=stream.readChar())!==Tags.TagClosebrace){__readRaw(stream,ostream,tag,binary);}
ostream.write(tag);}
function RawReader(stream,binary){defineProperties(this,{stream:{value:stream},binary:{value:!!binary,writable:true},readRaw:{value:function(){return readRaw(stream,this.binary);}}});}
global.hprose.RawReader=RawReader;var fakeReaderRefer=createObject(null,{set:{value:function(){}},read:{value:function(){unexpectedTag(Tags.TagRef);}},reset:{value:function(){}}});function RealReaderRefer(){defineProperties(this,{ref:{value:[]}});}
defineProperties(RealReaderRefer.prototype,{set:{value:function(val){this.ref.push(val);}},read:{value:function(index){return this.ref[index];}},reset:{value:function(){this.ref.length=0;}}});function realReaderRefer(){return new RealReaderRefer();}
function getter(str){var obj=global;var names=str.split('.');var i;for(i=0;i<names.length;i++){obj=obj[names[i]];if(obj===undefined){return null;}}
return obj;}
function findClass(cn,poslist,i,c){if(i<poslist.length){var pos=poslist[i];cn[pos]=c;var cls=findClass(cn,poslist,i+1,'.');if(i+1<poslist.length){if(cls===null){cls=findClass(cn,poslist,i+1,'_');}}
return cls;}
var classname=cn.join('');try{var cl=getter(classname);return((typeof(cl)==='function')?cl:null);}catch(e){return null;}}
function getClass(classname){var cls=ClassManager.getClass(classname);if(cls){return cls;}
cls=getter(classname);if(typeof(cls)==='function'){ClassManager.register(cls,classname);return cls;}
var poslist=[];var pos=classname.indexOf('_');while(pos>=0){poslist[poslist.length]=pos;pos=classname.indexOf('_',pos+1);}
if(poslist.length>0){var cn=classname.split('');cls=findClass(cn,poslist,0,'.');if(cls===null){cls=findClass(cn,poslist,0,'_');}
if(typeof(cls)==='function'){ClassManager.register(cls,classname);return cls;}}
cls=function(){};defineProperties(cls.prototype,{'getClassName':{value:function(){return classname;}}});ClassManager.register(cls,classname);return cls;}
function readInt(stream,tag){var s=stream.readUntil(tag);if(s.length===0){return 0;}
return parseInt(s,10);}
function unserialize(reader){var stream=reader.stream;var tag=stream.readChar();switch(tag){case'0':return 0;case'1':return 1;case'2':return 2;case'3':return 3;case'4':return 4;case'5':return 5;case'6':return 6;case'7':return 7;case'8':return 8;case'9':return 9;case Tags.TagInteger:return readIntegerWithoutTag(stream);case Tags.TagLong:return readLongWithoutTag(stream);case Tags.TagDouble:return readDoubleWithoutTag(stream);case Tags.TagNull:return null;case Tags.TagEmpty:return'';case Tags.TagTrue:return true;case Tags.TagFalse:return false;case Tags.TagNaN:return NaN;case Tags.TagInfinity:return readInfinityWithoutTag(stream);case Tags.TagDate:return readDateWithoutTag(reader);case Tags.TagTime:return readTimeWithoutTag(reader);case Tags.TagBytes:return readBinaryWithoutTag(reader);case Tags.TagUTF8Char:return readUTF8CharWithoutTag(reader);case Tags.TagString:return readStringWithoutTag(reader);case Tags.TagGuid:return readGuidWithoutTag(reader);case Tags.TagList:return readListWithoutTag(reader);case Tags.TagMap:return reader.useHarmonyMap?readHarmonyMapWithoutTag(reader):readMapWithoutTag(reader);case Tags.TagClass:readClass(reader);return readObject(reader);case Tags.TagObject:return readObjectWithoutTag(reader);case Tags.TagRef:return readRef(reader);case Tags.TagError:throw new Error(readString(reader));default:unexpectedTag(tag);}}
function readIntegerWithoutTag(stream){return readInt(stream,Tags.TagSemicolon);}
function readInteger(stream){var tag=stream.readChar();switch(tag){case'0':return 0;case'1':return 1;case'2':return 2;case'3':return 3;case'4':return 4;case'5':return 5;case'6':return 6;case'7':return 7;case'8':return 8;case'9':return 9;case Tags.TagInteger:return readIntegerWithoutTag(stream);default:unexpectedTag(tag);}}
function readLongWithoutTag(stream){var s=stream.readUntil(Tags.TagSemicolon);var l=parseInt(s,10);if(l.toString()===s){return l;}
return s;}
function readLong(stream){var tag=stream.readChar();switch(tag){case'0':return 0;case'1':return 1;case'2':return 2;case'3':return 3;case'4':return 4;case'5':return 5;case'6':return 6;case'7':return 7;case'8':return 8;case'9':return 9;case Tags.TagInteger:case Tags.TagLong:return readLongWithoutTag(stream);default:unexpectedTag(tag);}}
function readDoubleWithoutTag(stream){return parseFloat(stream.readUntil(Tags.TagSemicolon));}
function readDouble(stream){var tag=stream.readChar();switch(tag){case'0':return 0;case'1':return 1;case'2':return 2;case'3':return 3;case'4':return 4;case'5':return 5;case'6':return 6;case'7':return 7;case'8':return 8;case'9':return 9;case Tags.TagInteger:case Tags.TagLong:case Tags.TagDouble:return readDoubleWithoutTag(stream);case Tags.TagNaN:return NaN;case Tags.TagInfinity:return readInfinityWithoutTag(stream);default:unexpectedTag(tag);}}
function readInfinityWithoutTag(stream){return((stream.readChar()===Tags.TagNeg)?-Infinity:Infinity);}
function readBoolean(stream){var tag=stream.readChar();switch(tag){case Tags.TagTrue:return true;case Tags.TagFalse:return false;default:unexpectedTag(tag);}}
function readDateWithoutTag(reader){var stream=reader.stream;var year=parseInt(stream.read(4),10);var month=parseInt(stream.read(2),10)-1;var day=parseInt(stream.read(2),10);var date;var tag=stream.readChar();if(tag===Tags.TagTime){var hour=parseInt(stream.read(2),10);var minute=parseInt(stream.read(2),10);var second=parseInt(stream.read(2),10);var millisecond=0;tag=stream.readChar();if(tag===Tags.TagPoint){millisecond=parseInt(stream.read(3),10);tag=stream.readChar();if((tag>='0')&&(tag<='9')){stream.skip(2);tag=stream.readChar();if((tag>='0')&&(tag<='9')){stream.skip(2);tag=stream.readChar();}}}
if(tag===Tags.TagUTC){date=new Date(Date.UTC(year,month,day,hour,minute,second,millisecond));}
else{date=new Date(year,month,day,hour,minute,second,millisecond);}}
else if(tag===Tags.TagUTC){date=new Date(Date.UTC(year,month,day));}
else{date=new Date(year,month,day);}
reader.refer.set(date);return date;}
function readDate(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagDate:return readDateWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readTimeWithoutTag(reader){var stream=reader.stream;var time;var hour=parseInt(stream.read(2),10);var minute=parseInt(stream.read(2),10);var second=parseInt(stream.read(2),10);var millisecond=0;var tag=stream.readChar();if(tag===Tags.TagPoint){millisecond=parseInt(stream.read(3),10);tag=stream.readChar();if((tag>='0')&&(tag<='9')){stream.skip(2);tag=stream.readChar();if((tag>='0')&&(tag<='9')){stream.skip(2);tag=stream.readChar();}}}
if(tag===Tags.TagUTC){time=new Date(Date.UTC(1970,0,1,hour,minute,second,millisecond));}
else{time=new Date(1970,0,1,hour,minute,second,millisecond);}
reader.refer.set(time);return time;}
function readTime(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagTime:return readTimeWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readBinaryWithoutTag(reader){if(!reader.binary){throw new Error('The binary string does not support to unserialize in text mode.');}
var stream=reader.stream;var count=readInt(stream,Tags.TagQuote);var bs=new BinaryString(stream.read(count));stream.skip(1);reader.refer.set(bs);return bs;}
function readBinary(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagEmpty:return new BinaryString('');case Tags.TagBytes:return readBinaryWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readUTF8CharWithoutTag(reader){if(reader.binary){return reader.stream.readUTF8AsUTF16(1);}
return reader.stream.read(1);}
function _readString(reader){var stream=reader.stream;var count=readInt(stream,Tags.TagQuote);var s;if(reader.binary){s=stream.readUTF8AsUTF16(count);}
else{s=stream.read(count);}
stream.skip(1);return s;}
function readStringWithoutTag(reader){var s=_readString(reader);reader.refer.set(s);return s;}
function readString(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagEmpty:return'';case Tags.TagUTF8Char:return readUTF8CharWithoutTag(reader);case Tags.TagString:return readStringWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readGuidWithoutTag(reader){var stream=reader.stream;stream.skip(1);var s=stream.read(36);stream.skip(1);reader.refer.set(s);return s;}
function readGuid(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagGuid:return readGuidWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readListWithoutTag(reader){var stream=reader.stream;var list=[];reader.refer.set(list);var count=readInt(stream,Tags.TagOpenbrace);for(var i=0;i<count;i++){list[i]=unserialize(reader);}
stream.skip(1);return list;}
function readList(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagList:return readListWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readMapWithoutTag(reader){var stream=reader.stream;var map={};reader.refer.set(map);var count=readInt(stream,Tags.TagOpenbrace);for(var i=0;i<count;i++){var key=unserialize(reader);var value=unserialize(reader);map[key]=value;}
stream.skip(1);return map;}
function readMap(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagMap:return readMapWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readHarmonyMapWithoutTag(reader){var stream=reader.stream;var map=new Map();reader.refer.set(map);var count=readInt(stream,Tags.TagOpenbrace);for(var i=0;i<count;i++){var key=unserialize(reader);var value=unserialize(reader);map.set(key,value);}
stream.skip(1);return map;}
function readHarmonyMap(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagMap:return readHarmonyMapWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readObjectWithoutTag(reader){var stream=reader.stream;var cls=reader.classref[readInt(stream,Tags.TagOpenbrace)];var obj=new cls.classname();reader.refer.set(obj);for(var i=0;i<cls.count;i++){obj[cls.fields[i]]=unserialize(reader);}
stream.skip(1);return obj;}
function readObject(reader){var tag=reader.stream.readChar();switch(tag){case Tags.TagNull:return null;case Tags.TagClass:readClass(reader);return readObject(reader);case Tags.TagObject:return readObjectWithoutTag(reader);case Tags.TagRef:return readRef(reader);default:unexpectedTag(tag);}}
function readClass(reader){var stream=reader.stream;var classname=_readString(reader);var count=readInt(stream,Tags.TagOpenbrace);var fields=[];for(var i=0;i<count;i++){fields[i]=readString(reader);}
stream.skip(1);classname=getClass(classname);reader.classref.push({classname:classname,count:count,fields:fields});}
function readRef(reader){return reader.refer.read(readInt(reader.stream,Tags.TagSemicolon));}
function Reader(stream,simple,useHarmonyMap,binary){RawReader.call(this,stream,binary);this.useHarmonyMap=!!useHarmonyMap;defineProperties(this,{classref:{value:[]},refer:{value:simple?fakeReaderRefer:realReaderRefer()}});}
Reader.prototype=createObject(RawReader.prototype);Reader.prototype.constructor=Reader;defineProperties(Reader.prototype,{useHarmonyMap:{value:false,writable:true},checkTag:{value:function(expectTag,tag){if(tag===undefined){tag=this.stream.readChar();}
if(tag!==expectTag){unexpectedTag(tag,expectTag);}}},checkTags:{value:function(expectTags,tag){if(tag===undefined){tag=this.stream.readChar();}
if(expectTags.indexOf(tag)>=0){return tag;}
unexpectedTag(tag,expectTags);}},unserialize:{value:function(){return unserialize(this);}},readInteger:{value:function(){return readInteger(this.stream);}},readLong:{value:function(){return readLong(this.stream);}},readDouble:{value:function(){return readDouble(this.stream);}},readBoolean:{value:function(){return readBoolean(this.stream);}},readDateWithoutTag:{value:function(){return readDateWithoutTag(this);}},readDate:{value:function(){return readDate(this);}},readTimeWithoutTag:{value:function(){return readTimeWithoutTag(this);}},readTime:{value:function(){return readTime(this);}},readBinaryWithoutTag:{value:function(){return readBinaryWithoutTag(this);}},readBinary:{value:function(){return readBinary(this);}},readStringWithoutTag:{value:function(){return readStringWithoutTag(this);}},readString:{value:function(){return readString(this);}},readGuidWithoutTag:{value:function(){return readGuidWithoutTag(this);}},readGuid:{value:function(){return readGuid(this);}},readListWithoutTag:{value:function(){return readListWithoutTag(this);}},readList:{value:function(){return readList(this);}},readMapWithoutTag:{value:function(){return this.useHarmonyMap?readHarmonyMapWithoutTag(this):readMapWithoutTag(this);}},readMap:{value:function(){return this.useHarmonyMap?readHarmonyMap(this):readMap(this);}},readObjectWithoutTag:{value:function(){return readObjectWithoutTag(this);}},readObject:{value:function(){return readObject(this);}},reset:{value:function(){this.classref.length=0;this.refer.reset();}}});global.HproseReader=global.hprose.Reader=Reader;})(this||[eval][0]('this'));(function(global){'use strict';var StringIO=global.hprose.StringIO;var Writer=global.hprose.Writer;var Reader=global.hprose.Reader;var createObject=global.hprose.createObject;function serialize(value,simple,binary){var stream=new StringIO();var writer=new Writer(stream,simple,binary);writer.serialize(value);return stream.take();}
function unserialize(stream,simple,useHarmonyMap,binary){if(!(stream instanceof StringIO)){stream=new StringIO(stream);}
return new Reader(stream,simple,useHarmonyMap,binary).unserialize();}
global.hprose.Formatter=createObject(null,{serialize:{value:serialize},unserialize:{value:unserialize}});global.hprose.serialize=serialize;global.hprose.unserialize=unserialize;})(this||[eval][0]('this'));(function(global){'use strict';global.HproseResultMode=global.hprose.ResultMode={Normal:0,Serialized:1,Raw:2,RawWithEndTag:3};global.hprose.Normal=global.hprose.ResultMode.Normal;global.hprose.Serialized=global.hprose.ResultMode.Serialized;global.hprose.Raw=global.hprose.ResultMode.Raw;global.hprose.RawWithEndTag=global.hprose.ResultMode.RawWithEndTag;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var setImmediate=global.setImmediate;var Tags=global.hprose.Tags;var ResultMode=global.hprose.ResultMode;var StringIO=global.hprose.StringIO;var Writer=global.hprose.Writer;var Reader=global.hprose.Reader;var Future=global.hprose.Future;var defineProperties=global.hprose.defineProperties;var createObject=global.hprose.createObject;var parseuri=global.hprose.parseuri;var isObjectEmpty=global.hprose.isObjectEmpty;var GETFUNCTIONS=Tags.TagEnd;function noop(){}
var s_boolean='boolean';var s_string='string';var s_number='number';var s_function='function';var s_object='object';function Client(uri,functions,settings){var _uri,_uriList=[],_index=-1,_binary=false,_byref=false,_simple=false,_timeout=30000,_retry=10,_idempotent=false,_failswitch=false,_failround=0,_lock=false,_tasks=[],_useHarmonyMap=false,_filters=[],_batch=false,_batches=[],_ready=new Future(),_topics=createObject(null),_id=null,_keepAlive=true,_invokeHandler=invokeHandler,_batchInvokeHandler=batchInvokeHandler,_beforeFilterHandler=beforeFilterHandler,_afterFilterHandler=afterFilterHandler,_invokeHandlers=[],_batchInvokeHandlers=[],_beforeFilterHandlers=[],_afterFilterHandlers=[],self=this;function outputFilter(request,context){for(var i=0,n=_filters.length;i<n;i++){request=_filters[i].outputFilter(request,context);}
return request;}
function inputFilter(response,context){for(var i=_filters.length-1;i>=0;i--){response=_filters[i].inputFilter(response,context);}
return response;}
function beforeFilterHandler(request,context){request=outputFilter(request,context);return _afterFilterHandler(request,context).then(function(response){if(context.oneway){return;}
return inputFilter(response,context);});}
function afterFilterHandler(request,context){return self.sendAndReceive(request,context);}
function sendAndReceive(request,context,onsuccess,onerror){_beforeFilterHandler(request,context).then(onsuccess,function(e){if(retry(request,context,onsuccess,onerror)){return;}
onerror(e);});}
function failswitch(){var n=_uriList.length;if(n>1){var i=_index+1;if(i>=n){i=0;_failround++;}
_index=i;_uri=_uriList[_index];}
else{_failround++;}
if(typeof self.onfailswitch===s_function){self.onfailswitch(self);}}
function retry(data,context,onsuccess,onerror){if(context.failswitch){failswitch();}
if(context.idempotent&&(context.retried<context.retry)){var interval=++context.retried*500;if(context.failswitch){interval-=(_uriList.length-1)*500;}
if(interval>5000){interval=5000;}
if(interval>0){global.setTimeout(function(){sendAndReceive(data,context,onsuccess,onerror);},interval);}
else{sendAndReceive(data,context,onsuccess,onerror);}
return true;}
return false;}
function initService(stub){var context={retry:_retry,retried:0,idempotent:true,failswitch:true,timeout:_timeout,client:self,userdata:{}};var onsuccess=function(data){var error=null;try{var stream=new StringIO(data);var reader=new Reader(stream,true);var tag=stream.readChar();switch(tag){case Tags.TagError:error=new Error(reader.readString());break;case Tags.TagFunctions:var functions=reader.readList();reader.checkTag(Tags.TagEnd);setFunctions(stub,functions);break;default:error=new Error('Wrong Response:\r\n'+data);break;}}
catch(e){error=e;}
if(error!==null){_ready.reject(error);}
else{_ready.resolve(stub);}};sendAndReceive(GETFUNCTIONS,context,onsuccess,_ready.reject);}
function setFunction(stub,name){return function(){if(_batch){return _invoke(stub,name,Array.slice(arguments),true);}
else{return Future.all(arguments).then(function(args){return _invoke(stub,name,args,false);});}};}
function setMethods(stub,obj,namespace,name,methods){if(obj[name]!==undefined){return;}
obj[name]={};if(typeof(methods)===s_string||methods.constructor===Object){methods=[methods];}
if(Array.isArray(methods)){for(var i=0;i<methods.length;i++){var m=methods[i];if(typeof(m)===s_string){obj[name][m]=setFunction(stub,namespace+name+'_'+m);}
else{for(var n in m){setMethods(stub,obj[name],name+'_',n,m[n]);}}}}}
function setFunctions(stub,functions){for(var i=0;i<functions.length;i++){var f=functions[i];if(typeof(f)===s_string){if(stub[f]===undefined){stub[f]=setFunction(stub,f);}}
else{for(var name in f){setMethods(stub,stub,'',name,f[name]);}}}}
function copyargs(src,dest){var n=Math.min(src.length,dest.length);for(var i=0;i<n;++i){dest[i]=src[i];}}
function initContext(batch){if(batch){return{mode:ResultMode.Normal,binary:_binary,byref:_byref,simple:_simple,onsuccess:undefined,onerror:undefined,useHarmonyMap:_useHarmonyMap,client:self,userdata:{}};}
return{mode:ResultMode.Normal,binary:_binary,byref:_byref,simple:_simple,timeout:_timeout,retry:_retry,retried:0,idempotent:_idempotent,failswitch:_failswitch,oneway:false,sync:false,onsuccess:undefined,onerror:undefined,useHarmonyMap:_useHarmonyMap,client:self,userdata:{}};}
function getContext(stub,name,args,batch){var context=initContext(batch);if(name in stub){var method=stub[name];for(var key in method){if(key in context){context[key]=method[key];}}}
var i=0,n=args.length;for(;i<n;++i){if(typeof args[i]===s_function){break;}}
if(i===n){return context;}
var extra=args.splice(i,n-i);context.onsuccess=extra[0];n=extra.length;for(i=1;i<n;++i){var arg=extra[i];switch(typeof arg){case s_function:context.onerror=arg;break;case s_boolean:context.byref=arg;break;case s_number:context.mode=arg;break;case s_object:for(var k in arg){if(k in context){context[k]=arg[k];}}
break;}}
return context;}
function encode(name,args,context){var stream=new StringIO();stream.write(Tags.TagCall);var writer=new Writer(stream,context.simple,context.binary);writer.writeString(name);if(args.length>0||context.byref){writer.reset();writer.writeList(args);if(context.byref){writer.writeBoolean(true);}}
return stream;}
function __invoke(name,args,context,batch){if(_lock){return Future.promise(function(resolve,reject){_tasks.push({batch:batch,name:name,args:args,context:context,resolve:resolve,reject:reject});});}
if(batch){return multicall(name,args,context);}
return call(name,args,context);}
function _invoke(stub,name,args,batch){return __invoke(name,args,getContext(stub,name,args,batch),batch);}
function errorHandling(name,error,context,reject){try{if(context.onerror){context.onerror(name,error);}
else if(self.onerror){self.onerror(name,error);}
reject(error);}
catch(e){reject(e);}}
function invokeHandler(name,args,context){var request=encode(name,args,context);request.write(Tags.TagEnd);return Future.promise(function(resolve,reject){sendAndReceive(request.toString(),context,function(response){if(context.oneway){resolve();return;}
var result=null;var error=null;try{if(context.mode===ResultMode.RawWithEndTag){result=response;}
else if(context.mode===ResultMode.Raw){result=response.substring(0,response.length-1);}
else{var stream=new StringIO(response);var reader=new Reader(stream,false,context.useHarmonyMap,context.binary);var tag=stream.readChar();if(tag===Tags.TagResult){if(context.mode===ResultMode.Serialized){result=reader.readRaw();}
else{result=reader.unserialize();}
tag=stream.readChar();if(tag===Tags.TagArgument){reader.reset();var _args=reader.readList();copyargs(_args,args);tag=stream.readChar();}}
else if(tag===Tags.TagError){error=new Error(reader.readString());tag=stream.readChar();}
if(tag!==Tags.TagEnd){error=new Error('Wrong Response:\r\n'+response);}}}
catch(e){error=e;}
if(error){reject(error);}
else{resolve(result);}},reject);});}
function unlock(sync){return function(){if(sync){_lock=false;setImmediate(function(tasks){tasks.forEach(function(task){if('settings'in task){endBatch(task.settings).then(task.resolve,task.reject);}
else{__invoke(task.name,task.args,task.context,task.batch).then(task.resolve,task.reject);}});},_tasks);_tasks=[];}};}
function call(name,args,context){if(context.sync){_lock=true;}
var promise=Future.promise(function(resolve,reject){_invokeHandler(name,args,context).then(function(result){try{if(context.onsuccess){try{context.onsuccess(result,args);}
catch(e){if(context.onerror){context.onerror(name,e);}
reject(e);}}
resolve(result);}
catch(e){reject(e);}},function(error){errorHandling(name,error,context,reject);});});promise.whenComplete(unlock(context.sync));return promise;}
function multicall(name,args,context){return Future.promise(function(resolve,reject){_batches.push({args:args,name:name,context:context,resolve:resolve,reject:reject});});}
function getBatchContext(settings){var context={timeout:_timeout,binary:_binary,retry:_retry,retried:0,idempotent:_idempotent,failswitch:_failswitch,oneway:false,sync:false,client:self,userdata:{}};for(var k in settings){if(k in context){context[k]=settings[k];}}
return context;}
function batchInvokeHandler(batches,context){var request=batches.reduce(function(stream,item){item.context.binary=context.binary;stream.write(encode(item.name,item.args,item.context));return stream;},new StringIO());request.write(Tags.TagEnd);return Future.promise(function(resolve,reject){sendAndReceive(request.toString(),context,function(response){if(context.oneway){resolve(batches);return;}
var i=-1;var stream=new StringIO(response);var reader=new Reader(stream,false,false,context.binary);var tag=stream.readChar();try{while(tag!==Tags.TagEnd){var result=null;var error=null;var mode=batches[++i].context.mode;if(mode>=ResultMode.Raw){result=new StringIO();}
if(tag===Tags.TagResult){if(mode===ResultMode.Serialized){result=reader.readRaw();}
else if(mode>=ResultMode.Raw){result.write(Tags.TagResult);result.write(reader.readRaw());}
else{reader.useHarmonyMap=batches[i].context.useHarmonyMap;reader.reset();result=reader.unserialize();}
tag=stream.readChar();if(tag===Tags.TagArgument){if(mode>=ResultMode.Raw){result.write(Tags.TagArgument);result.write(reader.readRaw());}
else{reader.reset();var _args=reader.readList();copyargs(_args,batches[i].args);}
tag=stream.readChar();}}
else if(tag===Tags.TagError){if(mode>=ResultMode.Raw){result.write(Tags.TagError);result.write(reader.readRaw());}
else{reader.reset();error=new Error(reader.readString());}
tag=stream.readChar();}
if([Tags.TagEnd,Tags.TagResult,Tags.TagError].indexOf(tag)<0){reject(new Error('Wrong Response:\r\n'+response));return;}
if(mode>=ResultMode.Raw){if(mode===ResultMode.RawWithEndTag){result.write(Tags.TagEnd);}
batches[i].result=result.toString();}
else{batches[i].result=result;}
batches[i].error=error;}}
catch(e){reject(e);return;}
resolve(batches);},reject);});}
function beginBatch(){_batch=true;}
function endBatch(settings){settings=settings||{};_batch=false;if(_lock){return Future.promise(function(resolve,reject){_tasks.push({batch:true,settings:settings,resolve:resolve,reject:reject});});}
var batchSize=_batches.length;if(batchSize===0){return Future.value([]);}
var context=getBatchContext(settings);if(context.sync){_lock=true;}
var batches=_batches;_batches=[];var promise=Future.promise(function(resolve,reject){_batchInvokeHandler(batches,context).then(function(batches){batches.forEach(function(i){if(i.error){errorHandling(i.name,i.error,i.context,i.reject);}
else{try{if(i.context.onsuccess){try{i.context.onsuccess(i.result,i.args);}
catch(e){if(i.context.onerror){i.context.onerror(i.name,e);}
i.reject(e);}}
i.resolve(i.result);}
catch(e){i.reject(e);}}
delete i.context;delete i.resolve;delete i.reject;});resolve(batches);},function(error){batches.forEach(function(i){if('reject'in i){errorHandling(i.name,error,i.context,i.reject);}});reject(error);});});promise.whenComplete(unlock(context.sync));return promise;}
function getUri(){return _uri;}
function getUriList(){return _uriList;}
function setUriList(uriList){if(typeof(uriList)===s_string){_uriList=[uriList];}
else if(Array.isArray(uriList)){_uriList=uriList.slice(0);_uriList.sort(function(){return Math.random()-0.5;});}
else{return;}
_index=0;_uri=_uriList[_index];}
function getBinary(){return _binary;}
function setBinary(value){_binary=!!value;}
function getFailswitch(){return _failswitch;}
function setFailswitch(value){_failswitch=!!value;}
function getFailround(){return _failround;}
function getTimeout(){return _timeout;}
function setTimeout(value){if(typeof(value)==='number'){_timeout=value|0;}
else{_timeout=0;}}
function getRetry(){return _retry;}
function setRetry(value){if(typeof(value)==='number'){_retry=value|0;}
else{_retry=0;}}
function getIdempotent(){return _idempotent;}
function setIdempotent(value){_idempotent=!!value;}
function setKeepAlive(value){_keepAlive=!!value;}
function getKeepAlive(){return _keepAlive;}
function getByRef(){return _byref;}
function setByRef(value){_byref=!!value;}
function getSimpleMode(){return _simple;}
function setSimpleMode(value){_simple=!!value;}
function getUseHarmonyMap(){return _useHarmonyMap;}
function setUseHarmonyMap(value){_useHarmonyMap=!!value;}
function getFilter(){if(_filters.length===0){return null;}
if(_filters.length===1){return _filters[0];}
return _filters.slice();}
function setFilter(filter){_filters.length=0;if(Array.isArray(filter)){filter.forEach(function(filter){addFilter(filter);});}
else{addFilter(filter);}}
function addFilter(filter){if(filter&&typeof filter.inputFilter==='function'&&typeof filter.outputFilter==='function'){_filters.push(filter);}}
function removeFilter(filter){var i=_filters.indexOf(filter);if(i===-1){return false;}
_filters.splice(i,1);return true;}
function filters(){return _filters;}
function useService(uri,functions,create){if(create===undefined){if(typeof(functions)===s_boolean){create=functions;functions=false;}
if(!functions){if(typeof(uri)===s_boolean){create=uri;uri=false;}
else if(uri&&uri.constructor===Object||Array.isArray(uri)){functions=uri;uri=false;}}}
var stub=self;if(create){stub={};}
if(!uri&&!_uri){return new Error('You should set server uri first!');}
if(uri){_uri=uri;}
if(typeof(functions)===s_string||(functions&&functions.constructor===Object)){functions=[functions];}
if(!Array.isArray(functions)){setImmediate(initService,stub);return _ready;}
setFunctions(stub,functions);_ready.resolve(stub);return stub;}
function invoke(name,args,onsuccess){var argc=arguments.length;if((argc<1)||(typeof name!==s_string)){throw new Error('name must be a string');}
if(argc===1){args=[];}
if(argc===2){if(!Array.isArray(args)){var _args=[];if(typeof args!==s_function){_args.push(noop);}
_args.push(args);args=_args;}}
if(argc>2){if(typeof onsuccess!==s_function){args.push(noop);}
for(var i=2;i<argc;i++){args.push(arguments[i]);}}
return _invoke(self,name,args,_batch);}
function ready(onComplete,onError){return _ready.then(onComplete,onError);}
function getTopic(name,id){if(_topics[name]){var topics=_topics[name];if(topics[id]){return topics[id];}}
return null;}
function subscribe(name,id,callback,timeout,failswitch){if(typeof name!==s_string){throw new TypeError('topic name must be a string.');}
if(id===undefined||id===null){if(typeof callback===s_function){id=callback;}
else{throw new TypeError('callback must be a function.');}}
if(!_topics[name]){_topics[name]=createObject(null);}
if(typeof id===s_function){timeout=callback;callback=id;autoId().then(function(id){subscribe(name,id,callback,timeout,failswitch);});return;}
if(typeof callback!==s_function){throw new TypeError('callback must be a function.');}
if(Future.isPromise(id)){id.then(function(id){subscribe(name,id,callback,timeout,failswitch);});return;}
if(timeout===undefined){timeout=_timeout;}
var topic=getTopic(name,id);if(topic===null){var cb=function(){_invoke(self,name,[id,topic.handler,cb,{idempotent:true,failswitch:failswitch,timeout:timeout}],false);};topic={handler:function(result){var topic=getTopic(name,id);if(topic){if(result!==null){var callbacks=topic.callbacks;for(var i=0,n=callbacks.length;i<n;++i){try{callbacks[i](result);}
catch(e){}}}
if(getTopic(name,id)!==null){cb();}}},callbacks:[callback]};_topics[name][id]=topic;cb();}
else if(topic.callbacks.indexOf(callback)<0){topic.callbacks.push(callback);}}
function delTopic(topics,id,callback){if(topics){if(typeof callback===s_function){var topic=topics[id];if(topic){var callbacks=topic.callbacks;var p=callbacks.indexOf(callback);if(p>=0){callbacks[p]=callbacks[callbacks.length-1];callbacks.length--;}
if(callbacks.length===0){delete topics[id];}}}
else{delete topics[id];}}}
function unsubscribe(name,id,callback){if(typeof name!==s_string){throw new TypeError('topic name must be a string.');}
if(id===undefined||id===null){if(typeof callback===s_function){id=callback;}
else{delete _topics[name];return;}}
if(typeof id===s_function){callback=id;id=null;}
if(id===null){if(_id===null){if(_topics[name]){var topics=_topics[name];for(id in topics){delTopic(topics,id,callback);}}}
else{_id.then(function(id){unsubscribe(name,id,callback);});}}
else if(Future.isPromise(id)){id.then(function(id){unsubscribe(name,id,callback);});}
else{delTopic(_topics[name],id,callback);}
if(isObjectEmpty(_topics[name])){delete _topics[name];}}
function isSubscribed(name){return!!_topics[name];}
function subscribedList(){var list=[];for(var name in _topics){list.push(name);}
return list;}
function getId(){return _id;}
function autoId(){if(_id===null){_id=_invoke(self,'#',[],false);}
return _id;}
autoId.sync=true;autoId.idempotent=true;autoId.failswitch=true;function addInvokeHandler(handler){_invokeHandlers.push(handler);_invokeHandler=_invokeHandlers.reduceRight(function(next,handler){return function(name,args,context){return Future.sync(function(){return handler(name,args,context,next);});};},invokeHandler);}
function addBatchInvokeHandler(handler){_batchInvokeHandlers.push(handler);_batchInvokeHandler=_batchInvokeHandlers.reduceRight(function(next,handler){return function(batches,context){return Future.sync(function(){return handler(batches,context,next);});};},batchInvokeHandler);}
function addBeforeFilterHandler(handler){_beforeFilterHandlers.push(handler);_beforeFilterHandler=_beforeFilterHandlers.reduceRight(function(next,handler){return function(request,context){return Future.sync(function(){return handler(request,context,next);});};},beforeFilterHandler);}
function addAfterFilterHandler(handler){_afterFilterHandlers.push(handler);_afterFilterHandler=_afterFilterHandlers.reduceRight(function(next,handler){return function(request,context){return Future.sync(function(){return handler(request,context,next);});};},afterFilterHandler);}
function use(handler){addInvokeHandler(handler);return self;}
var batch=createObject(null,{begin:{value:beginBatch},end:{value:endBatch},use:{value:function(handler){addBatchInvokeHandler(handler);return batch;}}});var beforeFilter=createObject(null,{use:{value:function(handler){addBeforeFilterHandler(handler);return beforeFilter;}}});var afterFilter=createObject(null,{use:{value:function(handler){addAfterFilterHandler(handler);return afterFilter;}}});defineProperties(this,{'#':{value:autoId},onerror:{value:null,writable:true},onfailswitch:{value:null,writable:true},uri:{get:getUri},uriList:{get:getUriList,set:setUriList},id:{get:getId},binary:{get:getBinary,set:setBinary},failswitch:{get:getFailswitch,set:setFailswitch},failround:{get:getFailround},timeout:{get:getTimeout,set:setTimeout},retry:{get:getRetry,set:setRetry},idempotent:{get:getIdempotent,set:setIdempotent},keepAlive:{get:getKeepAlive,set:setKeepAlive},byref:{get:getByRef,set:setByRef},simple:{get:getSimpleMode,set:setSimpleMode},useHarmonyMap:{get:getUseHarmonyMap,set:setUseHarmonyMap},filter:{get:getFilter,set:setFilter},addFilter:{value:addFilter},removeFilter:{value:removeFilter},filters:{get:filters},useService:{value:useService},invoke:{value:invoke},ready:{value:ready},subscribe:{value:subscribe},unsubscribe:{value:unsubscribe},isSubscribed:{value:isSubscribed},subscribedList:{value:subscribedList},use:{value:use},batch:{value:batch},beforeFilter:{value:beforeFilter},afterFilter:{value:afterFilter}});{if((settings)&&(typeof settings===s_object)){['failswitch','timeout','retry','idempotent','keepAlive','byref','simple','useHarmonyMap','filter','binary'].forEach(function(key){if(key in settings){self[key](settings[key]);}});}
if(uri){setUriList(uri);useService(functions);}}}
function checkuri(uri){var parser=parseuri(uri);var protocol=parser.protocol;if(protocol==='http:'||protocol==='https:'||protocol==='tcp:'||protocol==='tcp4:'||protocol==='tcp6:'||protocol==='tcps:'||protocol==='tcp4s:'||protocol==='tcp6s:'||protocol==='tls:'||protocol==='ws:'||protocol==='wss:'){return;}
throw new Error('The '+protocol+' client isn\'t implemented.');}
function create(uri,functions,settings){try{return global.hprose.HttpClient.create(uri,functions,settings);}
catch(e){}
try{return global.hprose.TcpClient.create(uri,functions,settings);}
catch(e){}
try{return global.hprose.WebSocketClient.create(uri,functions,settings);}
catch(e){}
if(typeof uri==='string'){checkuri(uri);}
else if(Array.isArray(uri)){uri.forEach(function(uri){checkuri(uri);});throw new Error('Not support multiple protocol.');}
throw new Error('You should set server uri first!');}
defineProperties(Client,{create:{value:create}});global.HproseClient=global.hprose.Client=Client;})(this||[eval][0]('this'));(function(global){'use strict';if(typeof global.document==="undefined"){global.FlashHttpRequest={flashSupport:function(){return false;}};return;}
var document=global.document;var scripts=document.getElementsByTagName('script');var flashpath=scripts[scripts.length-1].getAttribute('flashpath')||'';scripts=null;var localfile=(global.location!==undefined&&global.location.protocol==='file:');var nativeXHR=(typeof(XMLHttpRequest)!=='undefined');var corsSupport=(!localfile&&nativeXHR&&'withCredentials'in new XMLHttpRequest());var flashID='flashhttprequest_as3';var flashSupport=false;var request=null;var callbackList=[];var jsTaskQueue=[];var swfTaskQueue=[];var jsReady=false;var swfReady=false;function checkFlash(){var flash='Shockwave Flash';var flashmime='application/x-shockwave-flash';var flashax='ShockwaveFlash.ShockwaveFlash';var plugins=navigator.plugins;var mimetypes=navigator.mimeTypes;var version=0;var ie=false;if(plugins&&plugins[flash]){version=plugins[flash].description;if(version&&!(mimetypes&&mimetypes[flashmime]&&!mimetypes[flashmime].enabledPlugin)){version=version.replace(/^.*\s+(\S+\s+\S+$)/,'$1');version=parseInt(version.replace(/^(.*)\..*$/,'$1'),10);}}
else if(global.ActiveXObject){try{ie=true;var ax=new global.ActiveXObject(flashax);if(ax){version=ax.GetVariable('$version');if(version){version=version.split(' ')[1].split(',');version=parseInt(version[0],10);}}}
catch(e){}}
if(version<10){return 0;}
else if(ie){return 1;}
else{return 2;}}
function setFlash(){var flashStatus=checkFlash();flashSupport=flashStatus>0;if(flashSupport){var div=document.createElement('div');div.style.width=0;div.style.height=0;if(flashStatus===1){div.innerHTML=['<object ','classid="clsid:D27CDB6E-AE6D-11cf-96B8-444553540000" ','type="application/x-shockwave-flash" ','width="0" height="0" id="',flashID,'" name="',flashID,'">','<param name="movie" value="',flashpath,'FlashHttpRequest.swf?',+(new Date()),'" />','<param name="allowScriptAccess" value="always" />','<param name="quality" value="high" />','<param name="wmode" value="opaque" />','</object>'].join('');}else{div.innerHTML='<embed id="'+flashID+'" '+'src="'+flashpath+'FlashHttpRequest.swf?'+(+(new Date()))+'" '+'type="application/x-shockwave-flash" '+'width="0" height="0" name="'+flashID+'" '+'allowScriptAccess="always" />';}
document.documentElement.appendChild(div);}}
function setJsReady(){if(jsReady){return;}
if(!localfile&&!corsSupport){setFlash();}
jsReady=true;while(jsTaskQueue.length>0){var task=jsTaskQueue.shift();if(typeof(task)==='function'){task();}}}
function post(url,header,data,callbackid,timeout,binary){data=encodeURIComponent(data);if(swfReady){request.post(url,header,data,callbackid,timeout,binary);}
else{swfTaskQueue.push(function(){request.post(url,header,data,callbackid,timeout,binary);});}}
var FlashHttpRequest={};FlashHttpRequest.flashSupport=function(){return flashSupport;};FlashHttpRequest.post=function(url,header,data,callback,timeout,binary){var callbackid=-1;if(callback){callbackid=callbackList.length;callbackList[callbackid]=callback;}
if(jsReady){post(url,header,data,callbackid,timeout,binary);}
else{jsTaskQueue.push(function(){post(url,header,data,callbackid,timeout,binary);});}};FlashHttpRequest.__callback=function(callbackid,data,error){data=(data!==null)?decodeURIComponent(data):null;error=(error!==null)?decodeURIComponent(error):null;if(typeof(callbackList[callbackid])==='function'){callbackList[callbackid](data,error);}
delete callbackList[callbackid];};FlashHttpRequest.__jsReady=function(){return jsReady;};FlashHttpRequest.__setSwfReady=function(){request=(navigator.appName.indexOf('Microsoft')!==-1)?global[flashID]:document[flashID];swfReady=true;global.__flash__removeCallback=function(instance,name){try{if(instance){instance[name]=null;}}
catch(flashEx){}};while(swfTaskQueue.length>0){var task=swfTaskQueue.shift();if(typeof(task)==='function'){task();}}};global.FlashHttpRequest=FlashHttpRequest;setJsReady();})(this||[eval][0]('this'));(function(global){'use strict';var parseuri=global.hprose.parseuri;var s_cookieManager={};function setCookie(headers,uri){var parser=parseuri(uri);var host=parser.host;var name,values;function _setCookie(value){var cookies,cookie,i;cookies=value.replace(/(^\s*)|(\s*$)/g,'').split(';');cookie={};value=cookies[0].replace(/(^\s*)|(\s*$)/g,'').split('=',2);if(value[1]===undefined){value[1]=null;}
cookie.name=value[0];cookie.value=value[1];for(i=1;i<cookies.length;i++){value=cookies[i].replace(/(^\s*)|(\s*$)/g,'').split('=',2);if(value[1]===undefined){value[1]=null;}
cookie[value[0].toUpperCase()]=value[1];}
if(cookie.PATH){if(cookie.PATH.charAt(0)==='"'){cookie.PATH=cookie.PATH.substr(1);}
if(cookie.PATH.charAt(cookie.PATH.length-1)==='"'){cookie.PATH=cookie.PATH.substr(0,cookie.PATH.length-1);}}
else{cookie.PATH='/';}
if(cookie.EXPIRES){cookie.EXPIRES=Date.parse(cookie.EXPIRES);}
if(cookie.DOMAIN){cookie.DOMAIN=cookie.DOMAIN.toLowerCase();}
else{cookie.DOMAIN=host;}
cookie.SECURE=(cookie.SECURE!==undefined);if(s_cookieManager[cookie.DOMAIN]===undefined){s_cookieManager[cookie.DOMAIN]={};}
s_cookieManager[cookie.DOMAIN][cookie.name]=cookie;}
for(name in headers){values=headers[name];name=name.toLowerCase();if((name==='set-cookie')||(name==='set-cookie2')){if(typeof(values)==='string'){values=[values];}
values.forEach(_setCookie);}}}
function getCookie(uri){var parser=parseuri(uri);var host=parser.host;var path=parser.path;var secure=(parser.protocol==='https:');var cookies=[];for(var domain in s_cookieManager){if(host.indexOf(domain)>-1){var names=[];for(var name in s_cookieManager[domain]){var cookie=s_cookieManager[domain][name];if(cookie.EXPIRES&&((new Date()).getTime()>cookie.EXPIRES)){names.push(name);}
else if(path.indexOf(cookie.PATH)===0){if(((secure&&cookie.SECURE)||!cookie.SECURE)&&(cookie.value!==null)){cookies.push(cookie.name+'='+cookie.value);}}}
for(var i in names){delete s_cookieManager[domain][names[i]];}}}
if(cookies.length>0){return cookies.join('; ');}
return'';}
global.hprose.cookieManager={setCookie:setCookie,getCookie:getCookie};})(this||[eval][0]('this'));(function(global,undefined){'use strict';var Client=global.hprose.Client;var Future=global.hprose.Future;var createObject=global.hprose.createObject;var defineProperties=global.hprose.defineProperties;var toBinaryString=global.hprose.toBinaryString;var toUint8Array=global.hprose.toUint8Array;var parseuri=global.hprose.parseuri;var cookieManager=global.hprose.cookieManager;var TimeoutError=global.TimeoutError;var FlashHttpRequest=global.FlashHttpRequest;var XMLHttpRequest=global.XMLHttpRequest;if(global.plus&&global.plus.net&&global.plus.net.XMLHttpRequest){XMLHttpRequest=global.plus.net.XMLHttpRequest;}
else if(global.document&&global.document.addEventListener){global.document.addEventListener("plusready",function(){XMLHttpRequest=global.plus.net.XMLHttpRequest;},false);}
var deviceone;try{deviceone=global.require("deviceone");}
catch(e){}
var localfile=(global.location!==undefined&&global.location.protocol==='file:');var nativeXHR=(typeof(XMLHttpRequest)!=='undefined');var corsSupport=(!localfile&&nativeXHR&&'withCredentials'in new XMLHttpRequest());var ActiveXObject=global.ActiveXObject;var XMLHttpNameCache=null;function createMSXMLHttp(){if(XMLHttpNameCache!==null){return new ActiveXObject(XMLHttpNameCache);}
var MSXML=['MSXML2.XMLHTTP','MSXML2.XMLHTTP.6.0','MSXML2.XMLHTTP.5.0','MSXML2.XMLHTTP.4.0','MSXML2.XMLHTTP.3.0','MsXML2.XMLHTTP.2.6','Microsoft.XMLHTTP','Microsoft.XMLHTTP.1.0','Microsoft.XMLHTTP.1'];var n=MSXML.length;for(var i=0;i<n;i++){try{var xhr=new ActiveXObject(MSXML[i]);XMLHttpNameCache=MSXML[i];return xhr;}
catch(e){}}
throw new Error('Could not find an installed XML parser');}
function createXHR(){if(nativeXHR){return new XMLHttpRequest();}
else if(ActiveXObject){return createMSXMLHttp();}
else{throw new Error("XMLHttpRequest is not supported by this browser.");}}
function noop(){}
if(nativeXHR&&typeof(Uint8Array)!=='undefined'&&!XMLHttpRequest.prototype.sendAsBinary){XMLHttpRequest.prototype.sendAsBinary=function(bs){var data=toUint8Array(bs);this.send(ArrayBuffer.isView?data:data.buffer);};}
function HttpClient(uri,functions,settings){if(this.constructor!==HttpClient){return new HttpClient(uri,functions,settings);}
Client.call(this,uri,functions,settings);var _header=createObject(null);var self=this;function xhrPost(request,env){var future=new Future();var xhr=createXHR();xhr.open('POST',self.uri(),true);if(corsSupport){xhr.withCredentials='true';}
for(var name in _header){xhr.setRequestHeader(name,_header[name]);}
if(!env.binary){xhr.setRequestHeader("Content-Type","text/plain; charset=UTF-8");}
xhr.onreadystatechange=function(){if(xhr.readyState===4){xhr.onreadystatechange=noop;if(xhr.status){if(xhr.status===200){if(env.binary){future.resolve(toBinaryString(xhr.response));}
else{future.resolve(xhr.responseText);}}
else{future.reject(new Error(xhr.status+':'+xhr.statusText));}}}};xhr.onerror=function(){future.reject(new Error('error'));};if(env.timeout>0){future=future.timeout(env.timeout).catchError(function(e){xhr.onreadystatechange=noop;xhr.onerror=noop;xhr.abort();throw e;},function(e){return e instanceof TimeoutError;});}
if(env.binary){xhr.responseType="arraybuffer";xhr.sendAsBinary(request);}
else{xhr.send(request);}
return future;}
function fhrPost(request,env){var future=new Future();var callback=function(data,error){if(error===null){future.resolve(data);}
else{future.reject(new Error(error));}};FlashHttpRequest.post(self.uri(),_header,request,callback,env.timeout,env.binary);return future;}
function apiPost(request,env){var future=new Future();var cookie=cookieManager.getCookie(self.uri());if(cookie!==''){_header['Cookie']=cookie;}
global.api.ajax({url:self.uri(),method:'post',data:{body:request},timeout:env.timeout,dataType:'text',headers:_header,returnAll:true,certificate:self.certificate},function(ret,err){if(ret){if(ret.statusCode===200){cookieManager.setCookie(ret.headers,self.uri());future.resolve(ret.body);}
else{future.reject(new Error(ret.statusCode+':'+ret.body));}}
else{future.reject(new Error(err.msg));}});return future;}
function deviceOnePost(request,env){var future=new Future();var http=deviceone.mm('do_Http');http.method="POST";http.timeout=env.timeout;http.contextType="text/plain; charset=UTF-8";http.url=self.uri();http.body=request;for(var name in _header){http.setRequestHeader(name,_header[name]);}
var cookie=cookieManager.getCookie(self.uri());if(cookie!==''){http.setRequestHeader('Cookie',cookie);}
http.on("success",function(data){var cookie=http.getResponseHeader('set-cookie');if(cookie){cookieManager.setCookie({'set-cookie':cookie},self.uri());}
future.resolve(data);});http.on("fail",function(result){future.reject(new Error(result.status+":"+result.data));});http.request();return future;}
function isCrossDomain(){if(global.location===undefined){return true;}
var parser=parseuri(self.uri());if(parser.protocol!==global.location.protocol){return true;}
if(parser.host!==global.location.host){return true;}
return false;}
function sendAndReceive(request,env){var fhr=(FlashHttpRequest.flashSupport()&&!localfile&&!corsSupport&&(env.binary||isCrossDomain()));var apicloud=(typeof(global.api)!=="undefined"&&typeof(global.api.ajax)!=="undefined");var future=fhr?fhrPost(request,env):apicloud?apiPost(request,env):deviceone?deviceOnePost(request,env):xhrPost(request,env);if(env.oneway){future.resolve();}
return future;}
function setHeader(name,value){if(name.toLowerCase()!=='content-type'){if(value){_header[name]=value;}
else{delete _header[name];}}}
defineProperties(this,{certificate:{value:null,writable:true},setHeader:{value:setHeader},sendAndReceive:{value:sendAndReceive}});}
function checkuri(uri){var parser=parseuri(uri);if(parser.protocol==='http:'||parser.protocol==='https:'){return;}
throw new Error('This client desn\'t support '+parser.protocol+' scheme.');}
function create(uri,functions,settings){if(typeof uri==='string'){checkuri(uri);}
else if(Array.isArray(uri)){uri.forEach(function(uri){checkuri(uri);});}
else{throw new Error('You should set server uri first!');}
return new HttpClient(uri,functions,settings);}
defineProperties(HttpClient,{create:{value:create}});global.HproseHttpClient=global.hprose.HttpClient=HttpClient;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var StringIO=global.hprose.StringIO;var Client=global.hprose.Client;var Future=global.hprose.Future;var TimeoutError=global.TimeoutError;var defineProperties=global.hprose.defineProperties;var toBinaryString=global.hprose.toBinaryString;var toUint8Array=global.hprose.toUint8Array;var parseuri=global.hprose.parseuri;var WebSocket=global.WebSocket||global.MozWebSocket;function noop(){}
function WebSocketClient(uri,functions,settings){if(typeof(WebSocket)==="undefined"){throw new Error('WebSocket is not supported by this browser.');}
if(this.constructor!==WebSocketClient){return new WebSocketClient(uri,functions,settings);}
Client.call(this,uri,functions,settings);var _id=0;var _count=0;var _futures=[];var _envs=[];var _requests=[];var _ready=null;var ws=null;var self=this;function getNextId(){return(_id<0x7fffffff)?++_id:_id=0;}
function send(id,request){var stream=new StringIO();stream.writeInt32BE(id);if(_envs[id].binary){stream.write(request);}
else{stream.writeUTF16AsUTF8(request);}
var message=toUint8Array(stream.take());if(ArrayBuffer.isView){ws.send(message);}
else{ws.send(message.buffer);}}
function onopen(e){_ready.resolve(e);}
function onmessage(e){var stream;if(typeof e.data==="string"){stream=new StringIO(StringIO.utf8Encode(e.data));}
else{stream=new StringIO(toBinaryString(e.data));}
var id=stream.readInt32BE();var future=_futures[id];var env=_envs[id];delete _futures[id];delete _envs[id];if(future!==undefined){--_count;var data=stream.read(stream.length()-4);if(!env.binary){data=StringIO.utf8Decode(data);}
future.resolve(data);}
if((_count<100)&&(_requests.length>0)){++_count;var request=_requests.pop();_ready.then(function(){send(request[0],request[1]);});}
if(_count===0&&!self.keepAlive()){close();}}
function onclose(e){_futures.forEach(function(future,id){future.reject(new Error(e.code+':'+e.reason));delete _futures[id];});_count=0;ws=null;}
function connect(){_ready=new Future();ws=new WebSocket(self.uri());ws.binaryType='arraybuffer';ws.onopen=onopen;ws.onmessage=onmessage;ws.onerror=noop;ws.onclose=onclose;}
function sendAndReceive(request,env){if(ws===null||ws.readyState===WebSocket.CLOSING||ws.readyState===WebSocket.CLOSED){connect();}
var future=new Future();var id=getNextId();_futures[id]=future;_envs[id]=env;if(env.timeout>0){future=future.timeout(env.timeout).catchError(function(e){delete _futures[id];--_count;throw e;},function(e){return e instanceof TimeoutError;});}
if(_count<100){++_count;_ready.then(function(){send(id,request);});}
else{_requests.push([id,request]);}
if(env.oneway){future.resolve();}
return future;}
function close(){if(ws!==null){ws.onopen=noop;ws.onmessage=noop;ws.onclose=noop;ws.close();}}
defineProperties(this,{sendAndReceive:{value:sendAndReceive},close:{value:close}});}
function checkuri(uri){var parser=parseuri(uri);if(parser.protocol==='ws:'||parser.protocol==='wss:'){return;}
throw new Error('This client desn\'t support '+parser.protocol+' scheme.');}
function create(uri,functions,settings){if(typeof uri==='string'){checkuri(uri);}
else if(Array.isArray(uri)){uri.forEach(function(uri){checkuri(uri);});}
else{throw new Error('You should set server uri first!');}
return new WebSocketClient(uri,functions,settings);}
defineProperties(WebSocketClient,{create:{value:create}});global.HproseWebSocketClient=global.hprose.WebSocketClient=WebSocketClient;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var Future=global.hprose.Future;var defineProperties=global.hprose.defineProperties;var toUint8Array=global.hprose.toUint8Array;var toBinaryString=global.hprose.toBinaryString;function noop(){}
var socketPool={};var socketManager=null;function receiveListener(info){var socket=socketPool[info.socketId];socket.onreceive(toBinaryString(info.data));}
function receiveErrorListener(info){var socket=socketPool[info.socketId];socket.onerror(info.resultCode);socket.destroy();}
function ChromeTcpSocket(){if(socketManager===null){socketManager=global.chrome.sockets.tcp;socketManager.onReceive.addListener(receiveListener);socketManager.onReceiveError.addListener(receiveErrorListener);}
this.socketId=new Future();this.connected=false;this.timeid=undefined;this.onclose=noop;this.onconnect=noop;this.onreceive=noop;this.onerror=noop;}
defineProperties(ChromeTcpSocket.prototype,{connect:{value:function(address,port,options){var self=this;socketManager.create({persistent:options&&options.persistent},function(createInfo){if(options){if('noDelay'in options){socketManager.setNoDelay(createInfo.socketId,options.noDelay,function(result){if(result<0){self.socketId.reject(result);socketManager.disconnect(createInfo.socketId);socketManager.close(createInfo.socketId);self.onclose();}});}
if('keepAlive'in options){socketManager.setKeepAlive(createInfo.socketId,options.keepAlive,function(result){if(result<0){self.socketId.reject(result);socketManager.disconnect(createInfo.socketId);socketManager.close(createInfo.socketId);self.onclose();}});}}
if(options&&options.tls){socketManager.setPaused(createInfo.socketId,true,function(){socketManager.connect(createInfo.socketId,address,port,function(result){if(result<0){self.socketId.reject(result);socketManager.disconnect(createInfo.socketId);socketManager.close(createInfo.socketId);self.onclose();}
else{socketManager.secure(createInfo.socketId,function(secureResult){if(secureResult!==0){self.socketId.reject(result);socketManager.disconnect(createInfo.socketId);socketManager.close(createInfo.socketId);self.onclose();}
else{socketManager.setPaused(createInfo.socketId,false,function(){self.socketId.resolve(createInfo.socketId);});}});}});});}
else{socketManager.connect(createInfo.socketId,address,port,function(result){if(result<0){self.socketId.reject(result);socketManager.disconnect(createInfo.socketId);socketManager.close(createInfo.socketId);self.onclose();}
else{self.socketId.resolve(createInfo.socketId);}});}});this.socketId.then(function(socketId){socketPool[socketId]=self;self.connected=true;self.onconnect(socketId);},function(reason){self.onerror(reason);});}},send:{value:function(data){data=toUint8Array(data).buffer;var self=this;var promise=new Future();this.socketId.then(function(socketId){socketManager.send(socketId,data,function(sendInfo){if(sendInfo.resultCode<0){self.onerror(sendInfo.resultCode);promise.reject(sendInfo.resultCode);self.destroy();}
else{promise.resolve(sendInfo.bytesSent);}});});return promise;}},destroy:{value:function(){var self=this;this.connected=false;this.socketId.then(function(socketId){socketManager.disconnect(socketId);socketManager.close(socketId);delete socketPool[socketId];self.onclose();});}},ref:{value:function(){this.socketId.then(function(socketId){socketManager.setPaused(socketId,false);});}},unref:{value:function(){this.socketId.then(function(socketId){socketManager.setPaused(socketId,true);});}},clearTimeout:{value:function(){if(this.timeid!==undefined){global.clearTimeout(this.timeid);}}},setTimeout:{value:function(timeout,fn){this.clearTimeout();this.timeid=global.setTimeout(fn,timeout);}}});global.hprose.ChromeTcpSocket=ChromeTcpSocket;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var Future=global.hprose.Future;var defineProperties=global.hprose.defineProperties;var atob=global.atob;var btoa=global.btoa;function noop(){}
var socketPool={};var socketManager=null;function APICloudTcpSocket(){if(socketManager===null){socketManager=global.api.require('socketManager');}
this.socketId=new Future();this.connected=false;this.timeid=undefined;this.onclose=noop;this.onconnect=noop;this.onreceive=noop;this.onerror=noop;}
defineProperties(APICloudTcpSocket.prototype,{connect:{value:function(address,port,options){var self=this;socketManager.createSocket({type:'tcp',host:address,port:port,timeout:options.timeout,returnBase64:true},function(ret){if(ret){switch(ret.state){case 101:break;case 102:self.socketId.resolve(ret.sid);break;case 103:self.onreceive(atob(ret.data.replace(/\s+/g,'')));break;case 201:self.socketId.reject(new Error('Create TCP socket failed'));break;case 202:self.socketId.reject(new Error('TCP connection failed'));break;case 203:self.onclose();self.onerror(new Error('Abnormal disconnect connection'));break;case 204:self.onclose();break;case 205:self.onclose();self.onerror(new Error('Unknown error'));break;}}});this.socketId.then(function(socketId){socketPool[socketId]=self;self.connected=true;self.onconnect(socketId);},function(reason){self.onerror(reason);});}},send:{value:function(data){var self=this;var promise=new Future();this.socketId.then(function(socketId){socketManager.write({sid:socketId,data:btoa(data),base64:true},function(ret,err){if(ret.status){promise.resolve();}
else{self.onerror(new Error(err.msg));promise.reject(err.msg);self.destroy();}});});return promise;}},destroy:{value:function(){var self=this;this.connected=false;this.socketId.then(function(socketId){socketManager.closeSocket({sid:socketId},function(ret,err){if(!ret.status){self.onerror(new Error(err.msg));}});delete socketPool[socketId];});}},ref:{value:noop},unref:{value:noop},clearTimeout:{value:function(){if(this.timeid!==undefined){global.clearTimeout(this.timeid);}}},setTimeout:{value:function(timeout,fn){this.clearTimeout();this.timeid=global.setTimeout(fn,timeout);}}});global.hprose.APICloudTcpSocket=APICloudTcpSocket;})(this||[eval][0]('this'));(function(global,undefined){'use strict';var ChromeTcpSocket=global.hprose.ChromeTcpSocket;var APICloudTcpSocket=global.hprose.APICloudTcpSocket;var Client=global.hprose.Client;var StringIO=global.hprose.StringIO;var Future=global.hprose.Future;var TimeoutError=global.TimeoutError;var createObject=global.hprose.createObject;var defineProperties=global.hprose.defineProperties;var parseuri=global.hprose.parseuri;function noop(){}
function setReceiveHandler(socket,onreceive){socket.onreceive=function(data){if(!('receiveEntry'in socket)){socket.receiveEntry={stream:new StringIO(),headerLength:4,dataLength:-1,id:null};}
var entry=socket.receiveEntry;var stream=entry.stream;var headerLength=entry.headerLength;var dataLength=entry.dataLength;var id=entry.id;stream.write(data);while(true){if((dataLength<0)&&(stream.length()>=headerLength)){dataLength=stream.readInt32BE();if((dataLength&0x80000000)!==0){dataLength&=0x7fffffff;headerLength=8;}}
if((headerLength===8)&&(id===null)&&(stream.length()>=headerLength)){id=stream.readInt32BE();}
if((dataLength>=0)&&((stream.length()-headerLength)>=dataLength)){onreceive(stream.read(dataLength),id);headerLength=4;id=null;stream.trunc();dataLength=-1;}
else{break;}}
entry.stream=stream;entry.headerLength=headerLength;entry.dataLength=dataLength;entry.id=id;};}
function TcpTransporter(client){if(client){this.client=client;this.uri=this.client.uri();this.size=0;this.pool=[];this.requests=[];}}
defineProperties(TcpTransporter.prototype,{create:{value:function(){var parser=parseuri(this.uri);var protocol=parser.protocol;var address=parser.hostname;var port=parseInt(parser.port,10);var tls;if(protocol==='tcp:'||protocol==='tcp4:'||protocol==='tcp6:'){tls=false;}
else if(protocol==='tcps:'||protocol==='tcp4s:'||protocol==='tcp6s:'||protocol==='tls:'){tls=true;}
else{throw new Error('Unsupported '+protocol+' protocol!');}
var conn;if(global.chrome&&global.chrome.sockets&&global.chrome.sockets.tcp){conn=new ChromeTcpSocket();}
else if(global.api&&global.api.require){conn=new APICloudTcpSocket();}
else{throw new Error('TCP Socket is not supported by this browser or platform.');}
var self=this;conn.connect(address,port,{persistent:true,tls:tls,timeout:this.client.timeout(),noDelay:this.client.noDelay(),keepAlive:this.client.keepAlive()});conn.onclose=function(){--self.size;};++this.size;return conn;}}});function FullDuplexTcpTransporter(client){TcpTransporter.call(this,client);}
FullDuplexTcpTransporter.prototype=createObject(TcpTransporter.prototype,{fetch:{value:function(){var pool=this.pool;while(pool.length>0){var conn=pool.pop();if(conn.connected){if(conn.count===0){conn.clearTimeout();conn.ref();}
return conn;}}
return null;}},init:{value:function(conn){var self=this;conn.count=0;conn.futures={};conn.envs={};conn.timeoutIds={};setReceiveHandler(conn,function(data,id){var future=conn.futures[id];var env=conn.envs[id];if(future){self.clean(conn,id);if(conn.count===0){self.recycle(conn);}
if(!env.binary){data=StringIO.utf8Decode(data);}
future.resolve(data);}});conn.onerror=function(e){var futures=conn.futures;for(var id in futures){var future=futures[id];self.clean(conn,id);future.reject(e);}};}},recycle:{value:function(conn){conn.unref();conn.setTimeout(this.client.poolTimeout(),function(){conn.destroy();});}},clean:{value:function(conn,id){if(conn.timeoutIds[id]!==undefined){global.clearTimeout(conn.timeoutIds[id]);delete conn.timeoutIds[id];}
delete conn.futures[id];delete conn.envs[id];--conn.count;this.sendNext(conn);}},sendNext:{value:function(conn){if(conn.count<10){if(this.requests.length>0){var request=this.requests.pop();request.push(conn);this.send.apply(this,request);}
else{if(this.pool.lastIndexOf(conn)<0){this.pool.push(conn);}}}}},send:{value:function(request,future,id,env,conn){var self=this;var timeout=env.timeout;if(timeout>0){conn.timeoutIds[id]=global.setTimeout(function(){self.clean(conn,id);if(conn.count===0){self.recycle(conn);}
future.reject(new TimeoutError('timeout'));},timeout);}
conn.count++;conn.futures[id]=future;conn.envs[id]=env;var len=request.length;var buf=new StringIO();buf.writeInt32BE(len|0x80000000);buf.writeInt32BE(id);if(env.binary){buf.write(request);}
else{buf.writeUTF16AsUTF8(request);}
conn.send(buf.take()).then(function(){self.sendNext(conn);});}},getNextId:{value:function(){return(this.nextid<0x7fffffff)?++this.nextid:this.nextid=0;}},sendAndReceive:{value:function(request,future,env){var conn=this.fetch();var id=this.getNextId();if(conn){this.send(request,future,id,env,conn);}
else if(this.size<this.client.maxPoolSize()){conn=this.create();conn.onerror=function(e){future.reject(e);};var self=this;conn.onconnect=function(){self.init(conn);self.send(request,future,id,env,conn);};}
else{this.requests.push([request,future,id,env]);}}}});FullDuplexTcpTransporter.prototype.constructor=TcpTransporter;function HalfDuplexTcpTransporter(client){TcpTransporter.call(this,client);}
HalfDuplexTcpTransporter.prototype=createObject(TcpTransporter.prototype,{fetch:{value:function(){var pool=this.pool;while(pool.length>0){var conn=pool.pop();if(conn.connected){conn.clearTimeout();conn.ref();return conn;}}
return null;}},recycle:{value:function(conn){if(this.pool.lastIndexOf(conn)<0){conn.unref();conn.setTimeout(this.client.poolTimeout(),function(){conn.destroy();});this.pool.push(conn);}}},clean:{value:function(conn){conn.onreceive=noop;conn.onerror=noop;if(conn.timeoutId!==undefined){global.clearTimeout(conn.timeoutId);delete conn.timeoutId;}}},sendNext:{value:function(conn){if(this.requests.length>0){var request=this.requests.pop();request.push(conn);this.send.apply(this,request);}
else{this.recycle(conn);}}},send:{value:function(request,future,env,conn){var self=this;var timeout=env.timeout;if(timeout>0){conn.timeoutId=global.setTimeout(function(){self.clean(conn);conn.destroy();future.reject(new TimeoutError('timeout'));},timeout);}
setReceiveHandler(conn,function(data){self.clean(conn);self.sendNext(conn);if(!env.binary){data=StringIO.utf8Decode(data);}
future.resolve(data);});conn.onerror=function(e){self.clean(conn);future.reject(e);};var len=request.length;var buf=new StringIO();buf.writeInt32BE(len);if(env.binary){buf.write(request);}
else{buf.writeUTF16AsUTF8(request);}
conn.send(buf.take());}},sendAndReceive:{value:function(request,future,env){var conn=this.fetch();if(conn){this.send(request,future,env,conn);}
else if(this.size<this.client.maxPoolSize()){conn=this.create();var self=this;conn.onerror=function(e){future.reject(e);};conn.onconnect=function(){self.send(request,future,env,conn);};}
else{this.requests.push([request,future,env]);}}}});HalfDuplexTcpTransporter.prototype.constructor=TcpTransporter;function TcpClient(uri,functions,settings){if(this.constructor!==TcpClient){return new TcpClient(uri,functions,settings);}
Client.call(this,uri,functions,settings);var self=this;var _noDelay=true;var _fullDuplex=false;var _maxPoolSize=10;var _poolTimeout=30000;var fdtrans=null;var hdtrans=null;function getNoDelay(){return _noDelay;}
function setNoDelay(value){_noDelay=!!value;}
function getFullDuplex(){return _fullDuplex;}
function setFullDuplex(value){_fullDuplex=!!value;}
function getMaxPoolSize(){return _maxPoolSize;}
function setMaxPoolSize(value){if(typeof(value)==='number'){_maxPoolSize=value|0;if(_maxPoolSize<1){_maxPoolSize=10;}}
else{_maxPoolSize=10;}}
function getPoolTimeout(){return _poolTimeout;}
function setPoolTimeout(value){if(typeof(value)==='number'){_poolTimeout=value|0;}
else{_poolTimeout=0;}}
function sendAndReceive(request,env){var future=new Future();if(_fullDuplex){if((fdtrans===null)||(fdtrans.uri!==self.uri)){fdtrans=new FullDuplexTcpTransporter(self);}
fdtrans.sendAndReceive(request,future,env);}
else{if((hdtrans===null)||(hdtrans.uri!==self.uri)){hdtrans=new HalfDuplexTcpTransporter(self);}
hdtrans.sendAndReceive(request,future,env);}
if(env.oneway){future.resolve();}
return future;}
defineProperties(this,{noDelay:{get:getNoDelay,set:setNoDelay},fullDuplex:{get:getFullDuplex,set:setFullDuplex},maxPoolSize:{get:getMaxPoolSize,set:setMaxPoolSize},poolTimeout:{get:getPoolTimeout,set:setPoolTimeout},sendAndReceive:{value:sendAndReceive}});}
function checkuri(uri){var parser=parseuri(uri);var protocol=parser.protocol;if(protocol==='tcp:'||protocol==='tcp4:'||protocol==='tcp6:'||protocol==='tcps:'||protocol==='tcp4s:'||protocol==='tcp6s:'||protocol==='tls:'){return;}
throw new Error('This client desn\'t support '+protocol+' scheme.');}
function create(uri,functions,settings){if(typeof uri==='string'){checkuri(uri);}
else if(Array.isArray(uri)){uri.forEach(function(uri){checkuri(uri);});}
else{throw new Error('You should set server uri first!');}
return new TcpClient(uri,functions,settings);}
defineProperties(TcpClient,{create:{value:create}});global.HproseTcpClient=global.hprose.TcpClient=TcpClient;})(this||[eval][0]('this'));(function(global){'use strict';var Tags=global.hprose.Tags;var StringIO=global.hprose.StringIO;var Writer=global.hprose.Writer;var Reader=global.hprose.Reader;var JSON=global.JSON;var s_id=1;function JSONRPCClientFilter(version){this.version=version||'2.0';}
JSONRPCClientFilter.prototype.inputFilter=function inputFilter(data){if(data.charAt(0)==='{'){data='['+data+']';}
var responses=JSON.parse(data);var stream=new StringIO();var writer=new Writer(stream,true);for(var i=0,n=responses.length;i<n;++i){var response=responses[i];if(response.error){stream.write(Tags.TagError);writer.writeString(response.error.message);}
else{stream.write(Tags.TagResult);writer.serialize(response.result);}}
stream.write(Tags.TagEnd);return stream.take();};JSONRPCClientFilter.prototype.outputFilter=function outputFilter(data){var requests=[];var stream=new StringIO(data);var reader=new Reader(stream,false,false);var tag=stream.readChar();do{var request={};if(tag===Tags.TagCall){request.method=reader.readString();tag=stream.readChar();if(tag===Tags.TagList){request.params=reader.readListWithoutTag();tag=stream.readChar();}
if(tag===Tags.TagTrue){tag=stream.readChar();}}
if(this.version==='1.1'){request.version='1.1';}
else if(this.version==='2.0'){request.jsonrpc='2.0';}
request.id=s_id++;requests.push(request);}while(tag===Tags.TagCall);if(requests.length>1){return JSON.stringify(requests);}
return JSON.stringify(requests[0]);};global.hprose.JSONRPCClientFilter=JSONRPCClientFilter;if(typeof(global.hprose.filter)==="undefined"){global.hprose.filter={JSONRPCClientFilter:global.hprose.JSONRPCClientFilter};}
else{global.hprose.filter.JSONRPCClientFilter=global.hprose.JSONRPCClientFilter;}})(this||[eval][0]('this'));(function(global){'use strict';global.hprose.common={Completer:global.hprose.Completer,Future:global.hprose.Future,ResultMode:global.hprose.ResultMode};global.hprose.io={StringIO:global.hprose.StringIO,ClassManager:global.hprose.ClassManager,Tags:global.hprose.Tags,RawReader:global.hprose.RawReader,Reader:global.hprose.Reader,Writer:global.hprose.Writer,Formatter:global.hprose.Formatter};global.hprose.client={Client:global.hprose.Client,HttpClient:global.hprose.HttpClient,TcpClient:global.hprose.TcpClient,WebSocketClient:global.hprose.WebSocketClient};if(typeof define==='function'){if(define.cmd){define('hprose',[],global.hprose);}
else if(define.amd){define('hprose',[],function(){return global.hprose;});}}
if(typeof module==='object'){module.exports=global.hprose;}})(this||[eval][0]('this'));