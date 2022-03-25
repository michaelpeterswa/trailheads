import Route from '@ember/routing/route';

export default class TrailheadsRoute extends Route {
    async model() {
        const response = await fetch('http://localhost:8080/api/trailheads', {method:'GET', 
        headers: {'Authorization': 'Basic ' + btoa('test@example.com:33a2e979-cad8-4318-adda-b33f266f17bc')}});
        const trailheads = await response.json();
        return trailheads;
    }
}
