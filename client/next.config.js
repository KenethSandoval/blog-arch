const removeImports = require('next-remove-imports')({});

module.exports = removeImports({
	reactStrictMode: true,
	images: {
        domains: ['images.unsplash.com'], // change domain images
		 		formats: ["image/webp"],
  },
})
