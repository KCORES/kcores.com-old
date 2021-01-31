// node.js
const mume = require("./mume");

async function main(markdownFile, htmlFile) {
    await mume.init();


    engine = new mume.MarkdownEngine({
        filePath: markdownFile,
        config: {
            previewTheme: "consolas.css",
            // revealjsTheme: "white.css"
            codeBlockTheme: "vs.css",
            printBackground: true,
            enableScriptExecution: true, // <= for running code chunks
        },
    });

    // html export
    await engine.htmlExport({ offline: false, runAllCodeChunks: true, htmlExportPath: htmlFile });



    return process.exit();
}

var arguments = process.argv.splice(2);
console.log(arguments)

main(arguments[0], arguments[1]);