import { defineStore } from 'pinia'

export const useBandStore = defineStore('band', {
    state: () => ({
        bands: [],
        loading: false,
    }),
    getters: {
        getBands(state) {
            return state.bands
        },
        getBand: (state) => (id) => {
            return state.bands.find((band) => band.id == id)
        },
        getLoading(state) {
            return state.loading
        }
    },
    actions: {
        async fetchBands() {
            this.loading = true;
            try {
                const response = await fetch('http://localhost/data');
                if (!response.ok) { // if HTTP-status is 404, 500, 403, etc 
                    throw new Error(`unable to fetch data, the API that provides data is down. Status: ${response.status}`); // reject the promise with an error message
                }
                const data = await response.json();
                this.bands = data;
                this.error = null; // successful request, clear previous error (if any)
            } catch (error) {
                this.error = error.message;
                console.error('Error fetching bands:', error);
            } finally {
                this.loading = false;
            }
        },
        async fetchBandById(id) {
            if (!this.getBands.length) {
                await this.fetchBands();
            }
            return this.getBand(id);
        }
    },
    persist: true,
})
