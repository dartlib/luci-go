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

import { BeforeEnterObserver } from '@vaadin/router';
import { customElement, html, LitElement, property } from 'lit-element';
import { consumeContext, provideContext } from '../libs/context';

const CLIENT_ID_KEY = 'client-id';

export const consumeClientId = consumeContext<'clientId', string>('clientId');
export const provideClientId = provideContext<'clientId', string>('clientId');
export const consumeResultDbHost = consumeContext<'resultDbHost', string>('resultDbHost');
export const provideResultDbHost = provideContext<'resultDbHost', string>('resultDbHost');

declare const CONFIGS: {
  resultDb: {
    host: string;
  },
};

// TODO(weiweilin): update description after clientId is moved to settings.js.
/**
 * Provides app configs to be shared across the app.
 * Loads the configs from local storage first and refreshes them later to avoid
 * blocking the rest of the page from rendering.
 */
export class AppConfigProviderElement extends LitElement implements BeforeEnterObserver {
  @property() clientId = '';
  @property() resultDbHost = '';

  async onBeforeEnter() {
    const cachedClientId = window.localStorage.getItem(CLIENT_ID_KEY);
    this.resultDbHost = CONFIGS.resultDb.host;
    if (cachedClientId === null) {
      await this.refreshClientId();
    } else {
      this.clientId = cachedClientId;
      this.refreshClientId();
    }
  }

  /**
   * Refreshes client ID and updates its cache.
   */
  private async refreshClientId() {
    const res = await fetch('/auth/api/v1/server/client_id');
    this.clientId = (await res.json()).client_id;
    window.localStorage.setItem(CLIENT_ID_KEY, this.clientId);
  }

  protected render() {
    return html`
      <slot></slot>
    `;
  }
}

customElement('tr-app-config-provider')(
  provideClientId(
    provideResultDbHost(
      AppConfigProviderElement,
    ),
  ),
);
