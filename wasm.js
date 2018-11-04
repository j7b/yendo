'use strict';

const WASM_URL = 'main.wasm';

var thingy;
var wasm;
var logLine = [];
var memory8;
var memoryInt32;

var cosa;

var td = new TextDecoder("utf-8");

var importObject = {
  env: {
    'main.whynot':function() {
      for (var i=0; i<arguments.length;i++) {
        console.log(arguments[i]);
      }
      var ptr = arguments[0];
      for (var i=0 ; i < arguments[1];i++) {
        console.log("ptr",ptr)
        console.log(memory8[ptr]);
        ptr+=4;
      }
      return arguments.length;
    },
    'main.update':function(ptr,len) {
      var sa = memory8.subarray(ptr,ptr+len);
      gl.clear(gl.COLOR_BUFFER_BIT);
			gl.bufferData(gl.ARRAY_BUFFER, sa, gl.DYNAMIC_DRAW);
			gl.drawArrays(gl.TRIANGLES, 0, 3);
    },
    'main.exec':function(ptr,len) {
      var ta = new Uint8Array(memory8.subarray(ptr,ptr+len));
      var s = td.decode(ta);
      var f = new Function(s);
      var a = Array.from(arguments).slice(2);
      f.apply(window,a);
    },
    'main.setstring':function(kptr,klen,vptr,vlen) {
      var dc = new TextDecoder("utf-8");
      var sa = memory8.subarray(kptr,kptr+klen);
      var key = dc.decode(sa);
      sa = memory8.subarray(vptr,vptr+vlen);
      var val = dc.decode(sa);
      Reflect.set(window,key,val)
    },
    'main.toString':function(ptr,len) {
      console.log("ptr " + ptr + " len "+len);
      var sa = memory8.subarray(ptr,ptr+len);
      console.dir(sa);
      cosa = new TextDecoder("utf-8").decode(sa);
    },
    'main.ClearColor':function(r,g,b,a) {
     gl.clearColor(r,g,b,a)
    },
    'main.Clear':function() {
      gl.clear(gl.COLOR_BUFFER_BIT);
    },
    'main.Set':function(k,v) {
      console.log("k "+ typeof k)
      console.dir(k);
      console.log("v"+typeof v)
      console.dir(v);
    },
    'main.requestAnimationFrame':function() {
      window.requestAnimationFrame(wasm.exports.frame);
    },
    io_get_stdout: function() {
      return 1;
    },
    resource_write: function(fd, ptr, len) {
      if (fd == 1) {
        for (let i=0; i<len; i++) {
          let c = memory8[ptr+i];
          if (c == 13) { // CR
            // ignore
          } else if (c == 10) { // LF
            // write line
            let line = new TextDecoder("utf-8").decode(new Uint8Array(logLine));
            logLine = [];
            console.log(line);
          } else {
            logLine.push(c);
          }
        }
      } else {
        console.error('invalid file descriptor:', fd);
      }
    },
  },
};

function init() {
  const derp = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
  };
  derp(fetch(WASM_URL), importObject).then(function(obj) {
    thingy = obj;
    wasm = obj.instance;
    memory8 = new Uint8Array(wasm.exports.memory.buffer);
    memoryInt32 = new Int32Array(wasm.exports.memory.buffer);
    wasm.exports.cwa_main();
  })
}

init();
