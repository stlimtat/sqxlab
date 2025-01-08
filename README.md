# sqxlab

Take home for sqxlab

## Headless Chrome based Disposable Browser and Orchestration with Linux User Segregation

## Objective

Running a full blown browser requires a Desktop environment and is heavy on resources. With a headless browser, it is possible to leverage the screencasting feature (often used for debugging purposes) to serve an interactive UI of the web page.

As the browser is the only interface accessible to the user, the minimal isolation can be achieved by segregating user access to corresponding linux users and launching the browser.

There are two parts of the problem:

### Part I: Setting up Chromium screencasting

1. Leveraging [CDP](https://chromedevtools.github.io/devtools-protocol/) to access screencasting while running chrome/chromium with a remote debugger.
1. Setting up a web application (screencast client) to connect to websocket and use the CDP for screencasting.
1. Accessing/Interacting with a single tab of the headless browser via the web application

### Part II: Orchestrating Part I for multiple users, i.e each user gets access to a dedicated browser session

The server will be running the following components:

1. An API server to manage linux users and the lifecycle of the chromium headless browser.
1. A proxy server to facilitate connection to the individual chromium websocket.
1. A chromium headless browser for each user session running/confined to dedicated linux users.
1. Optionally, a static or dynamic web application (screencast client) can be served from the server.

The client machine will:

1. Make an API call to the server to start a browser session.
1. The web application (either running locally or served via the server) will connect to the websocket connection and provide the interactable browser session.
