let showAll = false;
let chartInstance = null;

let primaryRequests = [];
let subRequests = [];

let primaryRequestCount = 0;
let primaryRequestCountEch = 0;
let primaryRequestCountPrivateDns = 0;

let subRequestCount = 0;
let subRequestCountEch = 0;
let subRequestCountPrivateDns = 0;

const chartLabels = [
  "Secure DNS",
  "Unsecure DNS",
  "ECH Enabled",
  "ECH Disabled",
];

const chartConfig = {
  type: "doughnut",
  data: {
    labels: chartLabels,
    datasets: [
      {
        label: "DNS",
        backgroundColor: ["#dab2ff", "#3c0366"],
        data: [0, 100],
      },
      {
        label: "ECH",
        backgroundColor: ["#ffa1ad", "#8b0836"],
        data: [0, 100],
      },
    ],
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: "bottom",
        labels: {
          generateLabels(chart) {
            const originalGenerate =
              Chart.overrides.doughnut.plugins.legend.labels.generateLabels;
            const labels = originalGenerate.call(this, chart);
            const flatColors = chart.data.datasets.flatMap(
              (ds) => ds.backgroundColor,
            );

            return labels.map((label) => {
              label.datasetIndex = Math.floor(label.index / 2);
              label.hidden = !chart.isDatasetVisible(label.datasetIndex);
              label.fillStyle = flatColors[label.index];
              return label;
            });
          },
        },
        onClick(e, legendItem, legend) {
          const chart = legend.chart;
          const datasetIndex = legendItem.datasetIndex;
          const meta = chart.getDatasetMeta(datasetIndex);

          meta.hidden =
            meta.hidden === null
              ? !chart.data.datasets[datasetIndex].hidden
              : null;
          chart.update();
        },
      },
      tooltip: {
        callbacks: {
          title(context) {
            const item = context?.[0];
            if (!item) return "";
            const labelIndex = item.datasetIndex * 2 + item.dataIndex;
            const label = item.chart.data.labels[labelIndex];
            return `${label}: ${item.formattedValue}%`;
          },
        },
      },
    },
  },
};

const setChartData = (total = 0, usedEch = 0, usedPrivateDns = 0) => {
  if (!chartInstance || total === 0) return;

  const toPercent = (value) => Math.round((value / total) * 100);

  const dnsPercent = toPercent(usedPrivateDns);
  const echPercent = toPercent(usedEch);

  chartInstance.data.datasets[0].data = [dnsPercent, 100 - dnsPercent];
  chartInstance.data.datasets[1].data = [echPercent, 100 - echPercent];

  chartInstance.update();
};

const updateChart = () => {
  const total = primaryRequestCount + (showAll ? subRequestCount : 0);
  const totalEch = primaryRequestCountEch + (showAll ? subRequestCountEch : 0);
  const totalDns =
    primaryRequestCountPrivateDns + (showAll ? subRequestCountPrivateDns : 0);

  setChartData(total || 100, totalEch, totalDns);
};

const updateNoRequestsMessage = () => {
  const noPrimary = primaryRequests.length === 0;
  const noSub = subRequests.length === 0;

  if (showAll && noPrimary && noSub) {
    $("#noRequests")
      .removeClass("hidden")
      .find("h2")
      .text("No requests found.");
  } else if (!showAll && noPrimary) {
    $("#noRequests")
      .removeClass("hidden")
      .find("h2")
      .text("No primary requests found.");
  } else {
    $("#noRequests").addClass("hidden");
  }
};

const addRequestCard = (data, isPrimary) => {
  const requestType = isPrimary ? "primaryRequest" : "subRequest";
  const typeClass = isPrimary ? "h-16" : "h-12 text-sm";
  const bgColor = isPrimary ? "bg-slate-900" : "bg-slate-700";
  const hiddenClass = isPrimary || showAll ? "" : "hidden";

  return `
    <div class="${hiddenClass} request" id="${requestType}-${data.requestInfo.requestId}">
      <div class="flex ${typeClass} w-full ${bgColor} items-center p-2 rounded-md mb-5">
        <div class="rounded-lg ${
          isPrimary ? "h-12 w-12" : "h-8 w-8"
        } bg-slate-300 flex items-center justify-center flex-shrink-0 text-slate-900 font-bold">
          <span>${data.requestInfo.statusCode}</span>
        </div>
        <div class="url ml-2 flex-grow overflow-hidden">
          <p class="text-ellipsis overflow-hidden whitespace-nowrap w-full text-white" title="${data.requestInfo.url}">
            ${data.requestInfo.url}
          </p>
        </div>
        <div class="stats flex items-center ml-2 flex-shrink-0 text-white">
          <span class="mr-1" title="DoH Usage">
            ${
              data.securityInfo.usedPrivateDns
                ? '<img src="/icons/check.svg" alt="✔️" class="w-6 h-6" />'
                : '<img src="/icons/error.svg" alt="❌" class="w-6 h-6" />'
            }
          </span>
          <span title="ECH Usage">
            ${
              data.securityInfo.usedEch
                ? '<img src="/icons/check.svg" alt="✔️" class="w-6 h-6" />'
                : '<img src="/icons/error.svg" alt="❌" class="w-6 h-6" />'
            }
          </span>
        </div>
      </div>
    </div>`;
};

const addPrimaryRequest = (data) => {
  primaryRequests.push(data);
  primaryRequestCount++;
  if (data.securityInfo.usedEch) primaryRequestCountEch++;
  if (data.securityInfo.usedPrivateDns) primaryRequestCountPrivateDns++;

  $("#requests").prepend(addRequestCard(data, true));
  updateChart();
  updateNoRequestsMessage();
};

const addSubRequest = (data) => {
  subRequests.push(data);
  subRequestCount++;
  if (data.securityInfo.usedEch) subRequestCountEch++;
  if (data.securityInfo.usedPrivateDns) subRequestCountPrivateDns++;

  $("#requests").prepend(addRequestCard(data, false));
  updateChart();
  updateNoRequestsMessage();
};

const exportData = async () => {
  let requests = [...primaryRequests, ...subRequests];

  if (!requests.length) {
    alert("No data to export.");
    return;
  }

  const blob = new Blob([JSON.stringify(requests, null, 2)], {
    type: "application/json",
  });
  const url = URL.createObjectURL(blob);

  await browser.downloads.download({
    url,
    filename: "doech_data.json",
    saveAs: true,
  });
};

const resetData = () => {
  primaryRequests = [];
  subRequests = [];
  primaryRequestCount = 0;
  primaryRequestCountEch = 0;
  primaryRequestCountPrivateDns = 0;
  subRequestCount = 0;
  subRequestCountEch = 0;
  subRequestCountPrivateDns = 0;

  $("#requests").empty();
  updateChart();
  updateNoRequestsMessage();

  browser.runtime.sendMessage({ type: "doech-reset" });
};

const showPrimaryRequests = () => {
  showAll = false;
  $("#showAllRequests").removeClass("bg-slate-700").addClass("bg-slate-900");
  $("#showPrimaryRequests")
    .removeClass("bg-slate-900")
    .addClass("bg-slate-700");
  $('#requests [id^="subRequest-"]').addClass("hidden");
  updateChart();
  updateNoRequestsMessage();
};

const showAllRequests = () => {
  showAll = true;
  $("#showPrimaryRequests")
    .removeClass("bg-slate-700")
    .addClass("bg-slate-900");
  $("#showAllRequests").removeClass("bg-slate-900").addClass("bg-slate-700");
  $('#requests [id^="subRequest-"]').removeClass("hidden");
  updateChart();
  updateNoRequestsMessage();
};

browser.runtime.onMessage.addListener(async (message) => {
  if (message.type !== "doech-update") return;

  const request = message.data;
  request.requestInfo.type === "main_frame"
    ? addPrimaryRequest(request)
    : addSubRequest(request);
});

document.addEventListener("DOMContentLoaded", async () => {
  const ctx = document.getElementById("chart");
  if (ctx) chartInstance = new Chart(ctx, chartConfig);

  const res = await browser.runtime.sendMessage({ type: "doech-init" });
  if (res && res.data) {
    res.data.forEach((data) =>
      data.requestInfo.type === "main_frame"
        ? addPrimaryRequest(data)
        : addSubRequest(data),
    );
  }

  $("#export").on("click", exportData);

  $("#showAllRequests").on("click", showAllRequests);

  $("#showPrimaryRequests").on("click", showPrimaryRequests);

  $("#reset").on("click", resetData);
});

window.addEventListener("resize", () => {
  chartInstance?.resize();
});
