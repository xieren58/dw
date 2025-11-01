import "./weapp-adapter/index.js";
import { EngineLoader } from './dmloader';
try {
    const fs = wx.getFileSystemManager();
    const filesToDelete = ['game.arci', 'game.projectc', 'game.arcd', 'game.dmanifest', 'game.public.der'];
    filesToDelete.forEach(file => {
        try {
            fs.unlinkSync(wx.env.USER_DATA_PATH + `/data/${file}`);
            console.log(`Deleted ${file}`);
        } catch (e) {
            console.log(`Failed to delete ${file}:`, e);
        }
    });

    EngineLoader.load("canvas", "{projectName}", function () {
        console.log('加载引擎...');
        require("{projectName}_wasm.js")
    });
} catch (error) {
    console.log(error)
}
