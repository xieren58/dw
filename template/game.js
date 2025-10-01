import "./weapp-adapter/index.js";
import { EngineLoader } from './dmloader';
try {
    EngineLoader.load("canvas", "{projectName}", function () {
        console.log('加载引擎...');
        require("{projectName}_wasm.js")
    });
} catch (error) {
    console.log(error)
}