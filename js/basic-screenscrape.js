// https://stackoverflow.com/questions/71437739/page-startscreencast-chrome-devtools-protocol-low-fps-issue
const puppeteer = require('puppeteer');
const fs = require('fs');
const path = require('path');

// Helper func
async function writeImageFilename(data) {
  const filename = path.join(
    '/tmp',
    Date.now().toString() + '.png'
  );
  fs.writeFileSync(filename, data, 'base64');
  return filename;
}

(async () => {
  const browser = await puppeteer.launch({
    headless: false,
  });
  const page = await browser.newPage();
  page.setViewport({
    width: 678,
    height: 1080,
  });

  await page.goto(`https://youtu.be/wxUo-AE5-24`,
    {
      timeout: 60000,
      waitUntil: 'networkidle0',
    }
  );

  const client = await page.target().createCDPSession();
  client.on('Page.screencastFrame', async (frameObject) => {
    await client.send('Page.screencastFrameAck', {
      sessionId: frameObject.sessionId,
    });
    await writeImageFilename(frameObject.data);
  });

  client.send('Page.startScreencast', {
    format: 'png',
    quality: 100,
    maxWidth: 678,
    maxHeight: 1080,
    everyNthFrame: 5,
  });
  await page.waitForNetworkIdle({ idleTime: 1000 });
  client.send('Page.stopScreencast');
  await browser.close();
})();
