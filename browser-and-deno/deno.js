import * as _ from './wasm_exec.js';
import { readAll } from 'https://deno.land/std@0.219.0/io/read_all.ts';
const go = new window.Go();
const f = await Deno.open('main.wasm');
const buf = await readAll(f);
const inst = await WebAssembly.instantiate(buf, go.importObject);
go.run(inst.instance);
