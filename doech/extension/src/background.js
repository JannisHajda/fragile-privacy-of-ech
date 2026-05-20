let requestsCache = null;

let storageInitPromise = browser.storage.local.get("requests").then((res) => {
  requestsCache = res.requests || [];
});

browser.webRequest.onHeadersReceived.addListener(
  (details) => receivedRequest(details),
  { urls: ["<all_urls>"] },
  ["blocking"],
);

const receivedRequest = async (details) => {
  const { requestId } = details;
  const securityInfo = await browser.webRequest.getSecurityInfo(requestId, {});
  const { usedEch, usedPrivateDns } = securityInfo;

  if (usedEch === undefined || usedPrivateDns === undefined) return;

  const data = {
    requestInfo: details,
    securityInfo: securityInfo,
  };

  await storageInitPromise;

  requestsCache.push(data);

  browser.storage.local.set({ requests: requestsCache });

  try {
    await browser.runtime.sendMessage({
      type: "doech-update",
      data,
    });
  } catch (e) {
    // sidebar is not opened, ignore the error
  }
};

browser.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (!message.type.startsWith("doech-")) return false;

  let messageType = message.type.replace("doech-", "");

  if (messageType === "init" || messageType === "export") {
    // 5. Send the memory cache directly to the sidebar
    return storageInitPromise.then(() => {
      return { data: requestsCache };
    });
  }

  if (messageType === "reset") {
    requestsCache = [];
    browser.storage.local.set({ requests: [] }).then(() => {
      sendResponse({ success: true });
    });
    return true;
  }
});
