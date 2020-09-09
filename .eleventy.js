const MarkdownIt = require("markdown-it");
const FrontMatter = require('markdown-it-front-matter');
const fs = require('fs');

module.exports = (eleventyConfig) => {
    const md = new MarkdownIt().use(FrontMatter, () => {});

    eleventyConfig.setTemplateFormats(["liquid", "md", "css"]);

    eleventyConfig.addShortcode("markdown", function (file) {
        const data = fs.readFileSync(file);
        return md.render(data.toString());
    });

    function sortByIndex(values) {
        let vals = [...values];
        return vals.sort((a,b) => Math.sign(a.data.index - b.data.index));
    }

    eleventyConfig.addFilter("sortByIndex", sortByIndex);

    return {
        dir: {
            input: "src"
        }
    };
};


