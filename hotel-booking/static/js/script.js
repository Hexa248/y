(function () {
  const cards = Array.from(document.querySelectorAll('.city-card'));
  const searchText = document.getElementById('searchText');
  const searchInfo = document.getElementById('searchInfo');

  function filterCards() {
    if (!searchText || cards.length === 0) return;

    const keyword = searchText.value.trim().toLowerCase();
    let visible = 0;

    cards.forEach((card) => {
      const city = (card.dataset.city || '').toLowerCase();
      const hotels = (card.dataset.hotels || '').toLowerCase();
      const match = keyword === '' || city.includes(keyword) || hotels.includes(keyword);
      card.style.display = match ? '' : 'none';
      if (match) visible++;
    });

    if (!searchInfo) return;
    searchInfo.textContent = keyword === ''
      ? 'Menampilkan semua kota.'
      : `Menampilkan ${visible} kota yang cocok dengan "${searchText.value}".`;
  }

  if (searchText) searchText.addEventListener('input', filterCards);

  const pickerBtn = document.getElementById('guestRoomPicker');
  const popover = document.getElementById('guestRoomPopover');
  const summary = document.getElementById('guestRoomSummary');
  const adultEl = document.getElementById('adultVal');
  const childEl = document.getElementById('childVal');
  const roomEl = document.getElementById('roomVal');

  let state = { adult: 2, child: 0, room: 1 };

  function renderGuestSummary() {
    if (!summary || !adultEl || !childEl || !roomEl) return;
    adultEl.textContent = String(state.adult);
    childEl.textContent = String(state.child);
    roomEl.textContent = String(state.room);
    summary.textContent = `${state.adult} Dewasa, ${state.child} Anak, ${state.room} Kamar`;
  }

  function updateCounter(key, step) {
    const min = key === 'adult' ? 1 : 0;
    if (key === 'room') {
      state.room = Math.max(1, Math.min(8, state.room + step));
    } else {
      state[key] = Math.max(min, Math.min(10, state[key] + step));
    }
    renderGuestSummary();
  }

  if (pickerBtn && popover) {
    pickerBtn.addEventListener('click', () => {
      popover.classList.toggle('hidden');
    });

    document.addEventListener('click', (e) => {
      if (!popover.contains(e.target) && !pickerBtn.contains(e.target)) {
        popover.classList.add('hidden');
      }
    });
  }

  document.querySelectorAll('[data-inc]').forEach((btn) => {
    btn.addEventListener('click', () => updateCounter(btn.dataset.inc, 1));
  });

  document.querySelectorAll('[data-dec]').forEach((btn) => {
    btn.addEventListener('click', () => updateCounter(btn.dataset.dec, -1));
  });

  renderGuestSummary();
})();
