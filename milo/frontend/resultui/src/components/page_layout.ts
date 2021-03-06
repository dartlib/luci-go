// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import '@chopsui/chops-signin';
import '@material/mwc-icon';
import { css, customElement, html, LitElement, property, PropertyValues } from 'lit-element';

import { consumeClientId } from '../context/config_provider';

/**
 * Renders page header, including a sign-in widget and a feedback button, at the
 * top of the child nodes.
 * Refreshes the page when a new clientId is provided.
 */
export class PageLayoutElement extends LitElement {
  @property() clientId!: string;

  private rendered = false;
  protected firstUpdated() {
    this.rendered = true;
  }

  protected shouldUpdate(changedProperties: PropertyValues) {
    if (this.rendered && changedProperties.has('clientId')) {
      // <chops-signin> (gapi.auth2) can not be initialized with a different
      // client-id. Refresh the page when a new clientId is provided.
      window.location.reload();
      return false;
    }
    return true;
  }

  protected render() {
    const feedbackComment = encodeURIComponent(
`From Link: ${document.location.href}
Please enter a description of the problem, with repro steps if applicable.
`);
    return html`
      <div id="container">
        <div id="title-container">
          <img id="chromium-icon" src="https://storage.googleapis.com/chrome-infra/lucy-small.png"/>
          <span id="headline">LUCI Test Results (BETA)</span>
        </div>
        <a
          id="feedback"
          title="Send Feedback"
          target="_blank"
          href="https://bugs.chromium.org/p/chromium/issues/entry?template=Build%20Infrastructure&components=Infra%3EPlatform%3EMilo%3EResultUI&labels=Pri-2&comment=${feedbackComment}"
        >
          <mwc-icon>feedback</mwc-icon>
        </a>
        <chops-signin id="signin" client-id=${this.clientId}></chops-signin>
      </div>
      <slot></slot>
    `;
  }

  static styles = css`
    :host {
      --header-height: 52px;
    }

    #container {
      box-sizing: border-box;
      height: var(--header-height);
      padding: 10px 0;
      display: flex;
    }
    #title-container {
      display: flex;
      flex: 1 1 100%;
      align-items: center;
      margin-left: 14px;
    }
    #chromium-icon {
      display: inline-block;
      width: 32px;
      height: 32px;
      margin-right: 8px;
    }
    #headline {
      color: rgb(95, 99, 104);
      font-family: "Google Sans", "Helvetica Neue", sans-serif;
      font-size: 18px;
      font-weight: 300;
      letter-spacing: 0.25px;
    }
    #signin {
      margin-right: 14px;
    }
    #feedback {
      height: 32px;
      width: 32px;
      --mdc-icon-size: 28px;
      margin-right: 14px;
      position: relative;
      color: black;
      opacity: 0.4;
    }
    #feedback>mwc-icon {
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
    }
    #feedback:hover {
      opacity: 0.6;
    }
  `;
}

customElement('tr-page-layout')(
  consumeClientId(
    PageLayoutElement,
  ),
);
