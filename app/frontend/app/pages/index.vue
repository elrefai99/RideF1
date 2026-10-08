<script setup lang="ts">
const sender = ref('hello@pbe.com')
const recipientInput = ref('')
const recipients = ref<string[]>([])
const subject = ref('')
const message = ref('')
const providerModalOpen = ref(false)
const notice = ref('')
const sendingProvider = ref('')
const config = useRuntimeConfig()

type DemoEmail = {
  from: string
  to: string[]
  subject: string
  message: string
}

type MockSendResponse = {
  message: string
}

onMounted(() => {
  void loadDemoEmail()
})

async function loadDemoEmail() {
  try {
    const demo = await $fetch<DemoEmail>(`${config.public.apiBase}/api/emails/demo`)
    sender.value = demo.from
    recipients.value = demo.to
    subject.value = demo.subject
    message.value = demo.message
  } catch {
    notice.value = 'Could not load the demo draft. Make sure the API is running.'
  }
}

function addRecipient() {
  const email = recipientInput.value.trim().replace(/,$/, '')
  if (!email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) return
  if (!recipients.value.includes(email)) recipients.value.push(email)
  recipientInput.value = ''
}

function removeRecipient(email: string) {
  recipients.value = recipients.value.filter((recipient) => recipient !== email)
}

function handleRecipientKeydown(event: KeyboardEvent) {
  if (event.key === 'Enter' || event.key === ',') {
    event.preventDefault()
    addRecipient()
  }
  if (event.key === 'Backspace' && !recipientInput.value && recipients.value.length) {
    recipients.value.pop()
  }
}

function focusRecipient(event: MouseEvent) {
  (event.currentTarget as HTMLElement).querySelector('input')?.focus()
}

function openProviderPicker() {
  addRecipient()
  if (!sender.value || !recipients.value.length || !subject.value.trim() || !message.value.trim()) {
    notice.value = 'Add a sender, recipient, subject, and message first.'
    return
  }
  notice.value = ''
  providerModalOpen.value = true
}

async function chooseProvider(provider: string) {
  sendingProvider.value = provider
  try {
    const result = await $fetch<MockSendResponse>(`${config.public.apiBase}/api/emails/send`, {
      method: 'POST',
      body: {
        from: sender.value,
        to: recipients.value,
        subject: subject.value,
        message: message.value,
        provider: provider.toLowerCase(),
      },
    })
    providerModalOpen.value = false
    notice.value = result.message
  } catch {
    notice.value = 'Could not reach the API. Check that it is running and try again.'
  } finally {
    sendingProvider.value = ''
  }
}

function closeModal(event: KeyboardEvent) {
  if (event.key === 'Escape') providerModalOpen.value = false
}

watch(providerModalOpen, (isOpen) => {
  if (import.meta.client) {
    if (isOpen) window.addEventListener('keydown', closeModal)
    else window.removeEventListener('keydown', closeModal)
  }
})
</script>

<template>
  <main class="workspace">
    <aside class="sidebar">
      <a class="brand" href="#" aria-label="PBE home">
        <span class="brand-mark">p</span>
        <span>pulse<span class="brand-light">beam</span></span>
      </a>

      <div class="workspace-label">WORKSPACE</div>
      <nav class="nav-list" aria-label="Main navigation">
        <a class="nav-item active" href="#compose">
          <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M4 5.5A1.5 1.5 0 0 1 5.5 4h13A1.5 1.5 0 0 1 20 5.5v13a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 4 18.5zM7 8h10M7 12h10M7 16h6" /></svg>
          <span>Compose</span>
        </a>
        <a class="nav-item" href="#inbox">
          <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M4 5h16v14H4zM4 13h4l1.5 2h5L16 13h4M8 9h8" /></svg>
          <span>Inbox</span>
          <span class="nav-count">8</span>
        </a>
        <a class="nav-item" href="#sent">
          <svg viewBox="0 0 24 24" aria-hidden="true"><path d="m21 3-7.5 18-3.8-7.7L2 9.5zM9.7 13.3 21 3" /></svg>
          <span>Sent</span>
        </a>
        <a class="nav-item" href="#drafts">
          <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M4 4h16v16H4zM8 9h8M8 13h8M8 17h4" /></svg>
          <span>Drafts</span>
          <span class="nav-count muted">2</span>
        </a>
      </nav>

      <div class="sidebar-bottom">
        <div class="plan-card">
          <span class="plan-icon">✦</span>
          <div>
            <strong>Make it personal</strong>
            <p>Emails that feel like you.</p>
          </div>
          <span class="plan-arrow">↗</span>
        </div>
        <button class="profile" type="button">
          <span class="avatar">MK</span>
          <span class="profile-copy"><strong>Mohamed Khaled</strong><small>Free workspace</small></span>
          <span class="profile-menu">···</span>
        </button>
      </div>
    </aside>

    <section id="compose" class="main-panel">
      <header class="topbar">
        <div class="breadcrumbs"><span>Workspace</span><span class="crumb-slash">/</span><strong>Compose</strong></div>
        <div class="topbar-actions">
          <span class="saved-status"><span class="saved-dot" /> All changes saved</span>
          <button class="help-button" type="button" aria-label="Help">?</button>
        </div>
      </header>

      <div class="page-content">
        <div class="page-heading">
          <div>
            <div class="eyebrow"><span class="eyebrow-line" /> YOUR NEXT GREAT EMAIL</div>
            <h1>Compose a message<span class="title-period">.</span></h1>
            <p class="subtitle">A thoughtful note can make someone's day.</p>
          </div>
          <button class="draft-button" type="button"><span>↗</span> Preview email</button>
        </div>

        <div class="composer-layout">
          <section class="composer-card" aria-label="Email composer">
            <div class="card-topline">
              <div class="compose-label"><span class="compose-icon">✎</span> New message</div>
              <button class="more-button" type="button" aria-label="More options">···</button>
            </div>

            <div class="field-row">
              <label for="sender">From</label>
              <div class="sender-control">
                <input id="sender" v-model="sender" type="email" autocomplete="email" placeholder="you@example.com">
                <span class="sender-chevron">⌄</span>
              </div>
            </div>

            <div class="field-row recipient-row">
              <label for="recipient">To</label>
              <div class="recipient-control" @click="focusRecipient">
                <span v-for="email in recipients" :key="email" class="recipient-chip">
                  {{ email }}
                  <button type="button" :aria-label="`Remove ${email}`" @click.stop="removeRecipient(email)">×</button>
                </span>
                <input
                  id="recipient"
                  v-model="recipientInput"
                  type="text"
                  autocomplete="off"
                  placeholder="Add email address…"
                  @keydown="handleRecipientKeydown"
                  @blur="addRecipient"
                >
              </div>
            </div>

            <div class="field-row">
              <label for="subject">Subject</label>
              <input id="subject" v-model="subject" class="plain-input" type="text" placeholder="Give your email a subject">
            </div>

            <div class="message-area">
              <div class="message-toolbar" aria-label="Formatting options">
                <button type="button" aria-label="Bold"><strong>B</strong></button>
                <button type="button" aria-label="Italic"><em>I</em></button>
                <button type="button" aria-label="Underline"><u>U</u></button>
                <span class="toolbar-divider" />
                <button type="button" aria-label="Add link">↗</button>
                <button type="button" aria-label="Add image">▧</button>
                <button type="button" aria-label="More formatting">···</button>
              </div>
              <textarea v-model="message" placeholder="Write something worth opening…" aria-label="Email message" />
              <div class="message-footer"><span>Keep it warm. Keep it human.</span><span>{{ message.length }} characters</span></div>
            </div>

            <div class="composer-footer">
              <div class="footer-tools">
                <button type="button" class="tool-button" aria-label="Attach file">↥</button>
                <button type="button" class="tool-button" aria-label="Add emoji">☺</button>
                <button type="button" class="tool-button" aria-label="Schedule email">◷</button>
              </div>
              <button class="send-button" type="button" @click="openProviderPicker">
                <span>Send email</span><span class="send-icon">➤</span>
              </button>
            </div>
          </section>

          <aside class="side-note">
            <div class="note-decoration"><span>✳</span><i /><i /><i /></div>
            <div class="note-tag">A LITTLE REMINDER</div>
            <h2>Write like you're talking to one person.</h2>
            <p>The best emails don't feel like emails. They feel like you took a moment to reach out.</p>
            <div class="note-bottom"><span class="note-line" /> <span>You've got this</span></div>
          </aside>
        </div>

        <p v-if="notice" class="notice" role="status">{{ notice }}</p>
        <footer class="page-footer"><span>Made for meaningful messages</span><span>✦ &nbsp;A little more human, every time.</span></footer>
      </div>
    </section>

    <div v-if="providerModalOpen" class="modal-backdrop" @click.self="providerModalOpen = false">
      <section class="provider-modal" role="dialog" aria-modal="true" aria-labelledby="provider-title">
        <button class="modal-close" type="button" aria-label="Close" @click="providerModalOpen = false">×</button>
        <div class="modal-mark">✉</div>
        <div class="modal-eyebrow">ONE LAST THING</div>
        <h2 id="provider-title">How would you like<br>to send this?</h2>
        <p>Choose the email service you'd like to use.</p>
        <button class="provider-option" type="button" :disabled="!!sendingProvider" @click="chooseProvider('Gmail')">
          <span class="gmail-logo"><svg viewBox="0 0 24 24" aria-hidden="true"><path fill="#4285f4" d="M1.8 6.5v11.2c0 .7.5 1.2 1.2 1.2h3.1V9.3z"/><path fill="#34a853" d="M17.9 18.9h3.1c.7 0 1.2-.5 1.2-1.2V6.5l-4.3 2.8z"/><path fill="#ea4335" d="m3 5.1 9 6.2 9-6.2c-.3-.4-.7-.7-1.2-.7H4.2c-.5 0-.9.3-1.2.7"/><path fill="#c5221f" d="M6.1 9.3V19h11.8V9.3L12 13.4z"/></svg></span>
          <span class="provider-copy"><strong>Gmail</strong><small>{{ sendingProvider === 'Gmail' ? 'Submitting demo request…' : 'Send from your Google account' }}</small></span>
          <span class="provider-arrow">→</span>
        </button>
        <button class="provider-option" type="button" :disabled="!!sendingProvider" @click="chooseProvider('Resend')">
          <span class="resend-logo">r</span>
          <span class="provider-copy"><strong>Resend</strong><small>{{ sendingProvider === 'Resend' ? 'Submitting demo request…' : 'Send with your Resend setup' }}</small></span>
          <span class="provider-arrow">→</span>
        </button>
        <div class="modal-footnote"><span>▣</span> Demo mode only. No email will be delivered.</div>
      </section>
    </div>
  </main>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=DM+Sans:wght@400;500;600;700&family=Manrope:wght@400;500;600;700;800&display=swap');

:root { font-family: 'DM Sans', sans-serif; color: #252821; background: #f7f8f5; font-synthesis: none; }
* { box-sizing: border-box; }
body { margin: 0; min-width: 320px; }
button, input, textarea { font: inherit; }
button { cursor: pointer; }
.workspace { min-height: 100vh; display: flex; }
.sidebar { width: 248px; flex: 0 0 248px; padding: 27px 16px 18px; background: #fff; border-right: 1px solid #e9ebe5; display: flex; flex-direction: column; }
.brand { display: flex; align-items: center; gap: 10px; margin: 0 0 54px 8px; color: #252821; text-decoration: none; font: 800 19px Manrope, sans-serif; letter-spacing: -.8px; }
.brand-mark { width: 29px; height: 29px; display: grid; place-items: center; border-radius: 9px; color: white; background: #687d52; font: 800 20px Manrope, sans-serif; transform: rotate(-5deg); }
.brand-light { color: #71825f; font-weight: 500; }
.workspace-label, .note-tag, .eyebrow, .modal-eyebrow { color: #989d91; font-size: 9px; font-weight: 700; letter-spacing: 1.5px; }
.workspace-label { padding-left: 12px; margin-bottom: 13px; }
.nav-list { display: grid; gap: 5px; }
.nav-item { min-height: 43px; display: flex; align-items: center; gap: 12px; padding: 0 12px; color: #777d72; border-radius: 9px; text-decoration: none; font-size: 12px; font-weight: 500; }
.nav-item svg { width: 17px; height: 17px; fill: none; stroke: currentColor; stroke-width: 1.6; stroke-linecap: round; stroke-linejoin: round; }
.nav-item.active { color: #536743; background: #f0f3ec; font-weight: 700; }
.nav-count { margin-left: auto; color: #7e8974; font-size: 10px; }
.nav-count.muted { color: #b0b4aa; }
.sidebar-bottom { margin-top: auto; }
.plan-card { padding: 14px 12px; display: flex; align-items: flex-start; gap: 10px; background: #f8f7ef; border: 1px solid #efeddd; border-radius: 11px; }
.plan-icon { width: 25px; height: 25px; flex: 0 0 25px; display: grid; place-items: center; border-radius: 8px; background: #e6e9d5; color: #718052; font-size: 14px; }
.plan-card strong { display: block; margin: 1px 0 4px; font: 700 10px Manrope, sans-serif; }
.plan-card p { margin: 0; color: #898d7f; font-size: 9px; }
.plan-arrow { margin-left: auto; color: #899276; font-size: 13px; }
.profile { width: 100%; display: flex; align-items: center; gap: 10px; margin-top: 18px; padding: 10px 3px 2px; border: 0; border-top: 1px solid #eceee8; background: none; text-align: left; }
.avatar { width: 31px; height: 31px; display: grid; place-items: center; border-radius: 50%; background: #e3e9dc; color: #596e4a; font: 700 9px Manrope, sans-serif; }
.profile-copy { display: grid; gap: 3px; }
.profile-copy strong { font-size: 10px; font-weight: 700; }
.profile-copy small { color: #989d93; font-size: 9px; }
.profile-menu { margin-left: auto; color: #93988d; font-size: 16px; }
.main-panel { flex: 1; min-width: 0; }
.topbar { height: 66px; display: flex; align-items: center; justify-content: space-between; padding: 0 54px; border-bottom: 1px solid #e9ebe5; background: rgba(255,255,255,.62); }
.breadcrumbs { display: flex; align-items: center; gap: 10px; color: #9a9e95; font-size: 11px; }
.crumb-slash { color: #c1c5bc; }
.breadcrumbs strong { color: #4e5549; font-weight: 600; }
.topbar-actions { display: flex; align-items: center; gap: 21px; }
.saved-status { display: flex; align-items: center; gap: 7px; color: #92988d; font-size: 10px; }
.saved-dot { width: 6px; height: 6px; border-radius: 50%; background: #89a379; }
.help-button { width: 23px; height: 23px; border: 1px solid #dfe2da; border-radius: 50%; color: #7d8378; background: white; font-size: 11px; }
.page-content { max-width: 1110px; margin: 0 auto; padding: 53px 54px 24px; }
.page-heading { display: flex; align-items: flex-end; justify-content: space-between; margin-bottom: 28px; }
.eyebrow { display: flex; align-items: center; gap: 8px; margin-bottom: 13px; color: #819070; font-size: 9px; }
.eyebrow-line { width: 18px; height: 1px; background: #829372; }
h1, h2 { font-family: Manrope, sans-serif; }
h1 { margin: 0; color: #252921; font-size: clamp(26px, 3vw, 35px); font-weight: 700; letter-spacing: -1.5px; }
.title-period { color: #8a9b72; }
.subtitle { margin: 8px 0 0; color: #92978d; font-size: 12px; }
.draft-button { padding: 9px 13px; color: #6c7564; border: 1px solid #e3e6de; border-radius: 7px; background: #fff; font-size: 10px; font-weight: 600; }
.draft-button span { margin-right: 5px; color: #78896b; }
.composer-layout { display: grid; grid-template-columns: minmax(0, 1fr) 238px; align-items: start; gap: 20px; }
.composer-card { overflow: hidden; background: white; border: 1px solid #e8eae4; border-radius: 12px; box-shadow: 0 5px 18px rgba(43,52,34,.025); }
.card-topline { height: 52px; display: flex; align-items: center; justify-content: space-between; padding: 0 21px; border-bottom: 1px solid #eff0ec; }
.compose-label { display: flex; align-items: center; gap: 9px; font: 700 12px Manrope, sans-serif; }
.compose-icon { width: 24px; height: 24px; display: grid; place-items: center; border-radius: 7px; color: #6f805e; background: #f1f4ed; font-size: 15px; }
.more-button { border: 0; color: #8a9085; background: none; font-size: 18px; }
.field-row { min-height: 57px; display: flex; align-items: center; gap: 18px; padding: 10px 21px; border-bottom: 1px solid #f0f1ed; }
.field-row > label { width: 47px; flex: 0 0 47px; color: #92978e; font-size: 11px; }
.sender-control { display: flex; align-items: center; gap: 10px; }
.sender-control input, .plain-input, .recipient-control input { min-width: 0; border: 0; outline: 0; color: #353a32; background: transparent; font-size: 11px; }
.sender-chevron { color: #8d9487; font-size: 14px; }
.recipient-row { align-items: flex-start; min-height: 58px; }
.recipient-row > label { padding-top: 8px; }
.recipient-control { min-height: 32px; flex: 1; display: flex; flex-wrap: wrap; align-items: center; gap: 6px; cursor: text; }
.recipient-chip { display: inline-flex; align-items: center; gap: 7px; padding: 5px 8px; border: 1px solid #e5e9df; border-radius: 6px; color: #59644e; background: #f6f8f3; font-size: 10px; }
.recipient-chip button { padding: 0; border: 0; color: #88917e; background: none; font-size: 14px; line-height: 10px; }
.recipient-control input { flex: 1; height: 28px; }
.plain-input { width: 100%; }
input::placeholder, textarea::placeholder { color: #b4b8b0; }
.message-area { min-height: 287px; padding: 13px 21px 0; }
.message-toolbar { height: 31px; display: flex; align-items: center; gap: 4px; }
.message-toolbar button { width: 26px; height: 25px; display: grid; place-items: center; border: 0; border-radius: 4px; color: #81877d; background: transparent; font-size: 11px; }
.message-toolbar button:hover { background: #f2f4ef; }
.toolbar-divider { width: 1px; height: 15px; margin: 0 5px; background: #e6e8e2; }
.message-area textarea { width: 100%; min-height: 205px; resize: vertical; padding: 15px 0; border: 0; outline: 0; color: #363a32; font-size: 12px; line-height: 1.8; }
.message-footer { display: flex; justify-content: space-between; padding: 8px 0 13px; color: #a4a89f; font-size: 9px; }
.composer-footer { height: 63px; display: flex; align-items: center; justify-content: space-between; padding: 0 21px; border-top: 1px solid #eff0ec; }
.footer-tools { display: flex; gap: 5px; }
.tool-button { width: 29px; height: 29px; border: 0; border-radius: 6px; color: #80877a; background: transparent; font-size: 15px; }
.tool-button:hover { background: #f4f5f1; }
.send-button { min-width: 136px; display: flex; align-items: center; justify-content: space-between; gap: 19px; padding: 10px 12px 10px 15px; border: 0; border-radius: 7px; color: white; background: #687d52; box-shadow: 0 3px 7px rgba(76,98,56,.16); font-size: 11px; font-weight: 600; }
.send-button:hover { background: #586c44; }
.send-icon { font-size: 11px; }
.side-note { position: relative; overflow: hidden; min-height: 254px; padding: 22px 20px; border: 1px solid #e8eadf; border-radius: 11px; background: #f1f2e9; }
.note-decoration { position: absolute; top: 7px; right: 9px; width: 76px; height: 68px; color: #d7dcc8; }
.note-decoration span { position: absolute; top: 11px; right: 15px; font-size: 36px; }
.note-decoration i { position: absolute; top: 25px; right: 34px; width: 25px; height: 1px; background: #d1d8c2; transform: rotate(49deg); }
.note-decoration i:nth-child(2) { transform: rotate(-49deg); }
.note-decoration i:nth-child(3) { top: 11px; right: 22px; width: 1px; height: 32px; transform: rotate(0); }
.note-tag { position: relative; z-index: 1; margin: 3px 0 18px; color: #82906f; font-size: 8px; }
.side-note h2 { max-width: 180px; margin: 0; color: #394331; font-size: 17px; font-weight: 700; line-height: 1.35; letter-spacing: -.5px; }
.side-note p { margin: 13px 0 25px; color: #818879; font-size: 10px; line-height: 1.7; }
.note-bottom { display: flex; align-items: center; gap: 8px; color: #778568; font: 600 9px Manrope, sans-serif; }
.note-line { width: 20px; height: 1px; background: #9ba78c; }
.notice { margin: 12px 0 0; color: #687d52; font-size: 11px; }
.page-footer { display: flex; justify-content: space-between; margin-top: 32px; color: #a1a59d; font-size: 9px; }
.modal-backdrop { position: fixed; z-index: 5; inset: 0; display: grid; place-items: center; padding: 20px; background: rgba(32,37,29,.38); backdrop-filter: blur(4px); }
.provider-modal { position: relative; width: min(100%, 410px); padding: 31px 32px 25px; border: 1px solid rgba(255,255,255,.8); border-radius: 16px; background: #fff; box-shadow: 0 24px 80px rgba(26,34,21,.2); text-align: center; animation: modal-in .18s ease-out; }
@keyframes modal-in { from { opacity: 0; transform: translateY(8px) scale(.985); } to { opacity: 1; transform: translateY(0) scale(1); } }
.modal-close { position: absolute; top: 13px; right: 15px; width: 28px; height: 28px; border: 0; border-radius: 50%; color: #8c9188; background: #f5f6f3; font-size: 19px; }
.modal-mark { width: 43px; height: 43px; display: grid; place-items: center; margin: 0 auto 15px; border-radius: 13px; color: #70815d; background: #f0f3eb; font-size: 20px; }
.modal-eyebrow { color: #89977a; font-size: 8px; }
.provider-modal h2 { margin: 10px 0 7px; color: #293025; font-size: 23px; line-height: 1.25; letter-spacing: -.7px; }
.provider-modal > p { margin: 0 0 22px; color: #92978e; font-size: 11px; }
.provider-option { width: 100%; display: flex; align-items: center; gap: 12px; margin-top: 9px; padding: 12px; border: 1px solid #e8ebe4; border-radius: 9px; background: #fff; text-align: left; transition: border-color .15s, background .15s, transform .15s; }
.provider-option:hover { transform: translateY(-1px); border-color: #cbd5c0; background: #fbfcfa; }
.gmail-logo, .resend-logo { width: 36px; height: 36px; display: grid; place-items: center; border: 1px solid #eff0ed; border-radius: 8px; background: #fff; }
.gmail-logo svg { width: 21px; height: 21px; }
.resend-logo { color: #151714; font: 700 22px Manrope, sans-serif; }
.provider-copy { display: grid; gap: 4px; }
.provider-copy strong { color: #343a30; font-size: 11px; }
.provider-copy small { color: #999e95; font-size: 9px; }
.provider-arrow { margin-left: auto; color: #8b9580; font-size: 16px; }
.modal-footnote { display: flex; align-items: center; justify-content: center; gap: 6px; margin-top: 20px; color: #a0a49c; font-size: 8px; }
.modal-footnote span { color: #849276; }
@media (max-width: 900px) { .sidebar { width: 205px; flex-basis: 205px; } .topbar { padding: 0 30px; } .page-content { padding: 42px 30px 20px; } .composer-layout { grid-template-columns: minmax(0, 1fr) 205px; gap: 14px; } }
@media (max-width: 700px) { .sidebar { width: 61px; flex-basis: 61px; padding: 20px 9px 13px; } .brand { margin: 0 auto 44px; } .brand > span:last-child, .workspace-label, .nav-item > span, .sidebar-bottom { display: none; } .nav-item { justify-content: center; padding: 0; } .topbar { height: 56px; padding: 0 17px; } .saved-status { font-size: 0; } .topbar-actions { gap: 11px; } .page-content { padding: 33px 17px 20px; } .page-heading { align-items: flex-start; gap: 12px; } .draft-button { flex: 0 0 auto; font-size: 0; } .draft-button span { font-size: 13px; } .composer-layout { display: block; } .side-note { display: none; } .page-footer { gap: 10px; font-size: 8px; } .page-footer span:last-child { text-align: right; } }
@media (max-width: 440px) { .brand-mark { width: 28px; height: 28px; } .page-heading { display: block; } .draft-button { margin-top: 17px; font-size: 10px; } .field-row { gap: 10px; padding-right: 13px; padding-left: 14px; } .field-row > label { width: 38px; flex-basis: 38px; } .card-topline, .composer-footer { padding-right: 14px; padding-left: 14px; } .message-area { padding-right: 14px; padding-left: 14px; } .message-footer { font-size: 8px; } .provider-modal { padding: 28px 20px 21px; } }
</style>
