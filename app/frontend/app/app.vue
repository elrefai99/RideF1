<script setup lang="ts">
type Sender = 'domain' | 'gmail'
type Email = { name: string; address: string; subject: string; preview: string; time: string; initials: string; color: string; unread?: boolean }

const section = ref<'send' | 'receive'>('send')
const sender = ref<Sender>('domain')
const recipients = ref<string[]>([])
const recipientInput = ref('')
const subject = ref('')
const message = ref('')
const messageEditor = ref<HTMLElement | null>(null)
const toast = ref('')
const inboxFilter = ref('')
const activeInboxTab = ref('All mail')
const sentMessages = ref<{ subject: string; recipientCount: number; time: string }[]>([])
const unreadOnly = computed(() => activeInboxTab.value === 'Unread')
let toastTimer: ReturnType<typeof setTimeout> | undefined

const emails = ref<Email[]>([
  { name: 'Jordan Davis', address: 'jordan@studio.co', subject: 'Re: Brand assets for the launch', preview: 'These look great! The updated files are exactly what we needed…', time: '10:24 AM', initials: 'JD', color: 'violet', unread: true },
  { name: 'Amira Lewis', address: 'amira@northstar.io', subject: 'Quick question about the proposal', preview: 'Hi Mohamed, I had one question about the timeline in section…', time: '9:15 AM', initials: 'AL', color: 'orange', unread: true },
  { name: 'Tom Kim', address: 'tom@fieldnotes.design', subject: 'Thanks for the introduction', preview: "Really appreciate you connecting us. I'll take it from here…", time: 'Yesterday', initials: 'TK', color: 'green' },
  { name: 'Notion', address: 'team@notion.so', subject: 'Your weekly workspace digest', preview: "Here's what happened in your workspace this week…", time: 'Oct 5', initials: 'NS', color: 'blue' },
])

const visibleEmails = computed(() => emails.value.filter((email) => {
  const matchesFilter = `${email.name} ${email.address} ${email.subject} ${email.preview}`.toLowerCase().includes(inboxFilter.value.toLowerCase())
  return matchesFilter && (!unreadOnly.value || email.unread)
}))

const senderAddress = computed(() => sender.value === 'gmail' ? 'Connect Gmail to choose an address' : 'hello@mydomain.com')
const wordCount = computed(() => message.value.trim() ? message.value.trim().split(/\s+/).length : 0)
const canSend = computed(() => recipients.value.length > 0 && subject.value.trim().length > 0)

function notify(text: string) {
  toast.value = text
  if (toastTimer) clearTimeout(toastTimer)
  toastTimer = setTimeout(() => { toast.value = '' }, 3200)
}

function addRecipient() {
  const address = recipientInput.value.trim().replace(/,$/, '')
  if (!address) return
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(address)) {
    notify('Enter a valid email address to add it.')
    return
  }
  if (!recipients.value.includes(address)) recipients.value.push(address)
  recipientInput.value = ''
}

function splitRecipients() {
  const entries = recipientInput.value.split(/[\s,;]+/).filter(Boolean)
  for (const entry of entries) {
    if (/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(entry) && !recipients.value.includes(entry)) recipients.value.push(entry)
  }
  recipientInput.value = ''
}

function handleRecipientKeydown(event: KeyboardEvent) {
  if (event.key === 'Enter' || event.key === ',') {
    event.preventDefault()
    addRecipient()
  } else if (event.key === 'Backspace' && !recipientInput.value && recipients.value.length) {
    removeRecipient(recipients.value[recipients.value.length - 1]!)
  }
}

function removeRecipient(address: string) {
  recipients.value = recipients.value.filter((item) => item !== address)
}

function changeSection(next: 'send' | 'receive') {
  section.value = next
}

function sendEmail() {
  if (!canSend.value) {
    notify('Add at least one recipient and a subject before sending.')
    return
  }
  sentMessages.value.unshift({ subject: subject.value.trim(), recipientCount: recipients.value.length, time: 'Just now' })
  subject.value = ''
  message.value = ''
  if (messageEditor.value) messageEditor.value.textContent = ''
  recipients.value = []
  notify('Demo email sent. No message was delivered.')
}

function saveDraft() {
  notify('Draft saved on this device.')
}

function toggleFormat(format: string) {
  if (import.meta.client) document.execCommand(format, false)
}

function markRead(email: Email) {
  email.unread = false
  notify(`Opened “${email.subject}”`)
}

onBeforeUnmount(() => { if (toastTimer) clearTimeout(toastTimer) })
</script>

<template>
  <div class="app-shell">
    <aside class="sidebar">
      <a class="brand" href="#" aria-label="Postmark home"><span class="brand-mark">p</span><span>postmark</span></a>
      <button class="workspace"><span class="workspace-icon">M</span><span class="workspace-copy"><b>Mohamed's workspace</b><small>Free plan</small></span><span class="chevrons">⌃<br>⌄</span></button>
      <div class="nav-label">WORKSPACE</div>
      <nav class="main-nav" aria-label="Main navigation">
        <button class="nav-item"><span class="nav-icon">▦</span>Overview</button>
        <button class="nav-item" :class="{ active: section === 'send' }" @click="changeSection('send')"><span class="nav-icon">↗</span>Send email</button>
        <button class="nav-item" :class="{ active: section === 'receive' }" @click="changeSection('receive')"><span class="nav-icon">↙</span>Receive email<span class="nav-count">3</span></button>
        <button class="nav-item"><span class="nav-icon">◷</span>Activity</button>
      </nav>
      <div class="nav-label connections-label">CONNECTIONS <button class="plus" aria-label="Add connection" @click="notify('Provider connections will be available when the mail service is configured.')">+</button></div>
      <button class="connection" @click="sender = 'domain'; changeSection('send')"><span class="connection-dot domain-dot">D</span><span><b>mydomain.com</b><small>Domain</small></span><span class="connected">●</span></button>
      <button class="connection" @click="sender = 'gmail'; changeSection('send'); notify('Gmail connection setup will be available soon.')"><span class="connection-dot gmail-dot">M</span><span><b>Gmail</b><small>Not connected</small></span><span class="connect-arrow">↗</span></button>
      <div class="sidebar-bottom"><div class="usage"><div><span>Monthly emails</span><b>248 <small>/ 1,000</small></b></div><div class="usage-track"><i /></div><small>Resets in 12 days</small></div><button class="user-card"><span class="avatar">MK</span><span><b>Mohamed Khaled</b><small>Personal account</small></span><span class="more">···</span></button></div>
    </aside>

    <main class="main-area">
      <header class="topbar"><div class="breadcrumbs"><span>Workspace</span><span class="slash">/</span><b>{{ section === 'send' ? 'Send email' : 'Received email' }}</b></div><div class="top-actions"><button class="help-button" @click="notify('Help center is coming soon.')">ⓘ <span>Help center</span></button><button class="icon-button" aria-label="Notifications" @click="notify('You’re all caught up.')">♧<i /></button><span class="avatar top-avatar">MK</span></div></header>

      <section v-if="section === 'send'" class="content">
        <div class="page-heading"><div><div class="eyebrow">EMAIL WORKSPACE <span class="live-dot" /> ALL SYSTEMS OPERATIONAL</div><h1>Send an email</h1><p>Write a message and send it from your connected address.</p></div><button class="secondary-button" @click="notify('You have 2 saved drafts.')"><span>▤</span> View drafts <span class="draft-count">2</span></button></div>
        <div class="composer-layout">
          <section class="compose-card">
            <div class="card-heading"><div><span class="step-number">01</span><h2>Compose your message</h2></div><span class="autosave"><i /> Draft saved</span></div>
            <form @submit.prevent="sendEmail">
              <div class="field-group"><label>FROM</label><div class="sender-row"><div class="sender-toggle" role="tablist" aria-label="Sender type"><button class="sender-option" :class="{ selected: sender === 'domain' }" type="button" @click="sender = 'domain'">⌂ &nbsp;Domain</button><button class="sender-option" :class="{ selected: sender === 'gmail' }" type="button" @click="sender = 'gmail'">◉ &nbsp;Gmail</button></div><div class="sender-select-wrap"><select :disabled="sender === 'gmail'" aria-label="From address"><option>{{ senderAddress }}</option><option v-if="sender === 'domain'">team@mydomain.com</option><option v-if="sender === 'domain'">support@mydomain.com</option></select><span class="select-caret">⌄</span></div></div><p class="field-hint">{{ sender === 'domain' ? 'Sending from your verified domain' : 'Connect a Gmail account to send from your inbox' }} <a v-if="sender === 'domain'" href="#" @click.prevent="notify('Domain settings are coming soon.')">Manage domain ↗</a></p></div>
              <div class="field-group"><label for="recipients">TO <span class="required">*</span></label><div class="recipient-box"><div class="recipient-chips"><span v-for="address in recipients" :key="address" class="recipient-chip">{{ address }}<button type="button" :aria-label="`Remove ${address}`" @click="removeRecipient(address)">×</button></span><input id="recipients" v-model="recipientInput" type="text" inputmode="email" placeholder="Add an email address…" autocomplete="off" @keydown="handleRecipientKeydown" @blur="splitRecipients"></div><button type="button" class="add-recipient" aria-label="Add recipient" @click="addRecipient">+</button></div><p class="field-hint">Press Enter or comma to add multiple recipients</p></div>
              <div class="field-group subject-group"><label for="subject">SUBJECT <span class="required">*</span></label><input id="subject" v-model="subject" class="text-input" maxlength="120" type="text" placeholder="A clear, thoughtful subject line"><span class="input-trailing">{{ subject.length }}/120</span></div>
              <div class="field-group message-group"><label for="message">MESSAGE</label><div class="editor"><div class="editor-toolbar"><button type="button" aria-label="Bold" @click="toggleFormat('bold')"><b>B</b></button><button type="button" aria-label="Italic" @click="toggleFormat('italic')"><i>I</i></button><button type="button" aria-label="Underline" @click="toggleFormat('underline')"><u>U</u></button><span class="toolbar-sep" /><button type="button" class="format-button">Paragraph⌄</button><span class="toolbar-sep" /><button type="button" aria-label="Bulleted list" @click="toggleFormat('insertUnorderedList')">☷</button><button type="button" aria-label="Insert link" @click="notify('Select text to add a link.')">↗</button><span class="toolbar-spacer" /><button type="button" title="Attach file" @click="notify('Attachments are coming soon.')">⌕</button></div><div id="message" ref="messageEditor" class="message-input" contenteditable="true" role="textbox" aria-label="Message" data-placeholder="Write your message here…" @input="message = ($event.target as HTMLElement).innerText"></div><div class="editor-footer"><span>Tip: keep it personal and concise</span><span>{{ wordCount }} {{ wordCount === 1 ? 'word' : 'words' }}</span></div></div></div>
              <div class="compose-footer"><span class="secure-note"><span>♧</span> Your message is encrypted and secure</span><div class="compose-actions"><button type="button" class="secondary-button" @click="saveDraft">Save draft</button><button type="submit" class="send-button"><span>↗</span> Send email</button></div></div>
            </form>
          </section>
          <aside class="right-column"><section class="side-card tips-card"><div class="side-card-heading"><span class="bulb">✦</span><h3>Before you send</h3></div><ul><li><span class="checkmark">✓</span><span><b>Verify your sender</b><small>{{ sender === 'domain' ? 'Domain is ready to send' : 'Connect your Gmail account' }}</small></span></li><li><span class="checkmark">✓</span><span><b>Check your recipients</b><small>Make sure every address is correct</small></span></li><li><span class="empty-check" /><span><b>Write a clear subject</b><small>Help people know what to expect</small></span></li></ul><a class="learn-link" href="#" @click.prevent="notify('Email best practices are coming soon.')">Email best practices <span>↗</span></a></section><section class="side-card activity-card"><div class="side-card-heading"><h3>Recent activity</h3><button aria-label="More activity options">···</button></div><div v-for="sent in sentMessages" :key="`${sent.subject}-${sent.time}`" class="activity-item"><span class="activity-icon sent-icon">↗</span><span><b>{{ sent.subject }}</b><small>Demo sent to {{ sent.recipientCount }} {{ sent.recipientCount === 1 ? 'person' : 'people' }} · {{ sent.time }}</small></span><span class="status-pill demo-pill">Demo</span></div><div class="activity-item"><span class="activity-icon sent-icon">↗</span><span><b>Sample: welcome to the team</b><small>Sample message · 10:42 AM</small></span><span class="status-pill demo-pill">Sample</span></div><div class="activity-item"><span class="activity-icon reply-icon">↙</span><span><b>Sample: project update</b><small>Received · Yesterday</small></span><span class="unread-dot" /></div><a class="learn-link" href="#" @click.prevent="changeSection('receive')">View all activity <span>→</span></a></section><div class="upgrade-card"><div class="sparkle">✳</div><b>Make more room for great ideas</b><p>Upgrade your plan for more emails and custom domains.</p><button @click="notify('Plan options are coming soon.')">Explore plans <span>→</span></button></div></aside>
        </div>
        <PageFooter />
      </section>

      <section v-else class="content receive-view"><div class="page-heading"><div><div class="eyebrow">DEMO INBOX <span class="live-dot" /> SAMPLE MESSAGES</div><h1>Received email</h1><p>Your incoming messages, all in one place.</p></div><button class="secondary-button" @click="notify('Sample inbox is up to date.')">↻ &nbsp; Refresh inbox</button></div><div class="inbox-card"><div class="inbox-toolbar"><div class="inbox-tabs"><button v-for="tab in ['All mail', 'Unread']" :key="tab" class="inbox-tab" :class="{ active: activeInboxTab === tab }" @click="activeInboxTab = tab">{{ tab }} <span>{{ tab === 'All mail' ? emails.length : emails.filter(email => email.unread).length }}</span></button></div><div class="inbox-search">⌕ <input v-model="inboxFilter" placeholder="Search your inbox"></div></div><div class="inbox-list"><button v-for="email in visibleEmails" :key="email.address" class="email-row" :class="{ unread: email.unread }" @click="markRead(email)"><span class="mail-avatar" :class="email.color">{{ email.initials }}</span><span class="email-info"><b>{{ email.name }} <small>{{ email.address }}</small></b><strong>{{ email.subject }}</strong><span>{{ email.preview }}</span></span><span class="email-meta"><time>{{ email.time }}</time><i v-if="email.unread" /></span></button><div v-if="visibleEmails.length === 0" class="empty-inbox">No messages match your search.</div></div><div class="inbox-bottom">Showing <b>{{ visibleEmails.length }}</b> of 24 messages <button @click="notify('You’re viewing the sample messages.')">Load more ↓</button></div></div><PageFooter /></section>
    </main>
    <Transition name="toast"><div v-if="toast" class="toast" role="status">{{ toast }}</div></Transition>
  </div>
</template>
