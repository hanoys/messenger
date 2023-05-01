const http = require('http');
const fs = require('fs');
const path = require('path');

const port = 8000;
const publicDir = path.join(__dirname, 'public');

const server = http.createServer((req, res) => {
    const filePath = path.join(publicDir, req.url === '/' ? 'index.html' : req.url);
    const fileExt = path.extname(filePath);
    const contentType = getContentType(fileExt);
    fs.readFile(filePath, (err, content) => {
        if (err) {
            res.writeHead(404, { 'Content-Type': 'text/plain' });
            res.end('404 Not Found');
        } else {
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content);
        }
    });
});

function getContentType(ext) {
    switch (ext) {
        case '.html':
            return 'text/html';
        case '.css':
            return 'text/css';
        case '.js':
            return 'text/javascript';
        default:
            return 'text/plain';
    }
}

server.listen(port, () => {
    console.log(`Server running at http://localhost:${port}/`);
});
