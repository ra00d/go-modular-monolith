const JsonViewerPlugin = (system) => ({
  wrapComponents: {
    HighlightCode: (Original) => (props) => {
      try {
        // console.log(system)
        const json = JSON.parse(props.children);
        // const jsonViewer = document.createElement("andypf-json-viewer")
        // jsonViewer.id = "json"
        // jsonViewer.expanded = 2
        // jsonViewer.indent = 2
        // jsonViewer.showDataTypes = true
        // jsonViewer.theme = "monokai"
        // jsonViewer.showToolbar = true
        // jsonViewer.showSize = true
        // jsonViewer.showCopy = true
        // jsonViewer.expandIconType = "square"
        // jsonViewer.expandEmpty = false
        // jsonViewer.data = json
        return system.React.createElement("andypf-json-viewer", {
          theme: "default-dark",
          data: JSON.stringify(json, null, 2),
          "show-toolbar": true,
          expanded: 2,
        });
      } catch (err) {
        console.error(err);
        return Original(props);
      }
    },
  },
});
function init(url) {
  window.onload = function () {
    window.ui = SwaggerUIBundle({
      url: url,
      dom_id: "#swagger-ui",
      deepLinking: true,
      presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
      plugins: [
        SwaggerUIBundle.plugins.DownloadUrl,
        SwaggerUIBundle.plugins.Auth,
        SwaggerUIBundle.plugins.RequestSnippets,
        JsonViewerPlugin,
      ],
      layout: "StandaloneLayout",
    });
  };
}
