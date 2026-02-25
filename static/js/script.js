function openPopup(title, img, desc, id) {
  const popup = document.getElementById('popup');
  if (!popup) return;
  document.getElementById('ptitle').innerText = title;
  document.getElementById('pimg').src = img;
  document.getElementById('pdesc').innerText = desc;
  document.getElementById('pdetail').href = '/hotel?id=' + id;
  popup.classList.remove('hidden');
}

function closePopup() {
  const popup = document.getElementById('popup');
  if (!popup) return;
  popup.classList.add('hidden');
}

function filterCards(selector, attrName, keyword) {
  const cards = document.querySelectorAll(selector);
  let visible = 0;
  cards.forEach((card) => {
    const name = (card.getAttribute(attrName) || '').toLowerCase();
    const show = !keyword || name.includes(keyword.toLowerCase());
    card.style.display = show ? '' : 'none';
    if (show) visible += 1;
  });
  return visible;
}

function closeAllSearchPopups(shell) {
  shell.querySelectorAll('.search-popup').forEach((p) => p.classList.add('hidden'));
}

function initSearchShell(shell) {
  const mode = shell.dataset.searchMode;
  const cityUrl = shell.dataset.cityUrl || '/city?name=';

  const destinationInput = shell.querySelector('.search-destination');
  const destinationPopup = shell.querySelector('.destination-popup');
  const optionButtons = shell.querySelectorAll('.search-option');
  const searchSubmit = shell.querySelector('.search-submit');

  const dateTrigger = shell.querySelector('.search-date-trigger');
  const guestTrigger = shell.querySelector('.search-guest-trigger');
  const datePopup = shell.querySelector('.date-popup');
  const guestPopup = shell.querySelector('.guest-popup');

  const checkinInput = shell.querySelector('.checkin-input');
  const checkoutInput = shell.querySelector('.checkout-input');
  const adultInput = shell.querySelector('.adult-input');
  const roomInput = shell.querySelector('.room-input');
  const dateSummary = shell.querySelector('.date-summary');
  const guestSummary = shell.querySelector('.guest-summary');

  function applyDestinationFilter() {
    const keyword = destinationInput.value.trim().toLowerCase();

    optionButtons.forEach((btn) => {
      const match = btn.innerText.toLowerCase().includes(keyword);
      btn.style.display = match ? '' : 'none';
    });

    if (mode === 'city') {
      filterCards('.city-tile', 'data-city-name', keyword);
    }
    if (mode === 'hotel') {
      const visible = filterCards('.hotel-card', 'data-hotel-name', keyword);
      const empty = document.getElementById('emptyHotelResult');
      if (empty) empty.classList.toggle('hidden', visible > 0);
    }
  }

  destinationInput.addEventListener('focus', () => {
    closeAllSearchPopups(shell);
    destinationPopup.classList.remove('hidden');
    applyDestinationFilter();
  });

  destinationInput.addEventListener('input', () => {
    destinationPopup.classList.remove('hidden');
    applyDestinationFilter();
  });

  destinationInput.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      searchSubmit.click();
    }
  });

  optionButtons.forEach((btn) => {
    btn.addEventListener('click', () => {
      destinationInput.value = btn.innerText.trim();
      destinationPopup.classList.add('hidden');
      applyDestinationFilter();
      if (mode === 'city') {
        window.location.href = cityUrl + encodeURIComponent(destinationInput.value.trim());
      }
    });
  });

  dateTrigger?.addEventListener('click', () => {
    closeAllSearchPopups(shell);
    datePopup.classList.remove('hidden');
  });

  guestTrigger?.addEventListener('click', () => {
    closeAllSearchPopups(shell);
    guestPopup.classList.remove('hidden');
  });

  shell.querySelector('.apply-date')?.addEventListener('click', () => {
    const checkin = checkinInput?.value;
    const checkout = checkoutInput?.value;
    if (checkin && checkout) {
      dateSummary.textContent = `${checkin} - ${checkout}`;
    }
    datePopup.classList.add('hidden');
  });

  shell.querySelector('.apply-guest')?.addEventListener('click', () => {
    const adults = adultInput?.value || 2;
    const rooms = roomInput?.value || 1;
    guestSummary.textContent = `${adults} Dewasa, ${rooms} Kamar`;
    guestPopup.classList.add('hidden');
  });

  searchSubmit?.addEventListener('click', () => {
    const keyword = destinationInput.value.trim();
    if (!keyword) return;

    if (mode === 'city') {
      const exact = Array.from(optionButtons).find((btn) => btn.innerText.trim().toLowerCase() === keyword.toLowerCase());
      const candidate = exact || Array.from(optionButtons).find((btn) => btn.style.display !== 'none');
      if (candidate) {
        window.location.href = cityUrl + encodeURIComponent(candidate.innerText.trim());
      }
      return;
    }

    applyDestinationFilter();
    const firstVisible = document.querySelector('.hotel-card:not([style*="display: none"])');
    if (firstVisible) {
      firstVisible.scrollIntoView({ behavior: 'smooth', block: 'center' });
    }
  });

  document.addEventListener('click', (e) => {
    if (!shell.contains(e.target)) {
      closeAllSearchPopups(shell);
    }
  });
}

document.querySelectorAll('.search-shell').forEach(initSearchShell);

document.addEventListener('click', (e) => {
  const popup = document.getElementById('popup');
  if (popup && !popup.classList.contains('hidden') && e.target.id === 'popup') closePopup();
});

document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape') {
    closePopup();
    document.querySelectorAll('.search-shell').forEach((shell) => closeAllSearchPopups(shell));
  }
});


function initDockedSearchAndNavbar() {
  const topbar = document.querySelector('.topbar');
  const shell = document.querySelector('.search-shell');
  if (!topbar && !shell) return;

  let lastY = window.scrollY;
  let shellTop = 0;
  let shellHeight = 0;
  let spacer = null;

  if (shell) {
    spacer = document.createElement('div');
    spacer.className = 'search-shell-placeholder';
    shell.parentNode.insertBefore(spacer, shell);
  }

  function recalc() {
    if (!shell) return;
    shell.classList.remove('is-docked');
    spacer.classList.remove('active');
    spacer.style.height = '0px';
    const rect = shell.getBoundingClientRect();
    shellTop = rect.top + window.scrollY;
    shellHeight = rect.height;
  }

  function onScroll() {
    const y = window.scrollY;

    if (topbar) {
      if (y > 90 && y > lastY + 2) {
        topbar.classList.add('topbar-hidden');
      } else {
        topbar.classList.remove('topbar-hidden');
      }
    }

    if (shell) {
      if (y >= shellTop - 6) {
        if (!shell.classList.contains('is-docked')) {
          shell.classList.add('is-docked');
          spacer.classList.add('active');
          spacer.style.height = `${shellHeight}px`;
        }
      } else {
        shell.classList.remove('is-docked');
        spacer.classList.remove('active');
        spacer.style.height = '0px';
      }
    }

    lastY = y;
  }

  recalc();
  onScroll();
  window.addEventListener('scroll', onScroll, { passive: true });
  window.addEventListener('resize', () => {
    recalc();
    onScroll();
  });
}

initDockedSearchAndNavbar();
