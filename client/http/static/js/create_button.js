function createButton(pageEl, fetchFunction) {
  const btn = document.createElement("button");

  btn.style.marginTop = "10px";
  btn.style.padding = "10px";
  btn.style.backgroundColor = "#000000";
  btn.style.color = "#ffffff";

  btn.innerText = "Load More...";

  pageEl.appendChild(btn);

  const pageInput = document.getElementById("page");

  btn.addEventListener("click", async function () {
    pageInput.value++;
    btn.remove();

    try {
      await fetchFunction();
    } catch (err) {
      console.log(err);
    }
  });
}
