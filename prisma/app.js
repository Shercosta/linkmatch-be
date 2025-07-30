const ResumeParser = require('pdf-resume-parser');
const fs = require('fs');
const path = require('path');

const filePath = process.argv[2];
if (!filePath) {
    console.error("❌ Please provide the PDF path: node app.js ./cv/yourfile.pdf");
    process.exit(1);
}

const parser = new ResumeParser.ResumeParser.ResumeParser();

async function parseResume(filePath) {
    const data = await parser.parseResume(filePath);
    const filename = path.basename(filePath, path.extname(filePath));

    const outputDir = path.join(__dirname, 'parsed-cv');
    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir, { recursive: true });
    }

    if (data) {
        const jsonData = JSON.stringify(data, null, 2);
        const outputPath = path.join(outputDir, `${filename}.json`);
        fs.writeFileSync(outputPath, jsonData);
        console.log(`✅ Resume parsed and saved to ${outputPath}`);
    }

    return data;
}

parseResume(filePath);
