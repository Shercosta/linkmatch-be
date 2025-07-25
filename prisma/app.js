const ResumeParser = require('pdf-resume-parser');

const parser = new ResumeParser.ResumeParser.ResumeParser();

async function parseResume(path) {
    const data = await parser.parseResume(path);
    // console.log(data)
    const fs = require('fs');

    if (data) {
        const jsonData = JSON.stringify(data, null, 2);
        fs.writeFileSync('data.json', jsonData);
    }
    return data
}

parseResume('./cv/april.pdf')
