const MarkdownIt = require("markdown-it");
const FrontMatter = require('markdown-it-front-matter');
const fs = require('fs');

module.exports = (eleventyConfig) => {
    const md = new MarkdownIt().use(FrontMatter, () => {});

    eleventyConfig.addShortcode("markdown", function (file) {
        const data = fs.readFileSync(file);
        return md.render(data.toString());
    });

    return {
        dir: {
            input: "src"
        }
    }
};


