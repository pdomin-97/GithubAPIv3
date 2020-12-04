import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FrontPageComponent } from '../front-page/front-page.component';
import { ResumeComponent } from '../resume/resume.component';

const routes: Routes = [
    {
        path: '',
        component: FrontPageComponent

    },
    {
        path: 'resume',
        component: ResumeComponent 
    },
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes)
    ],
    exports: [
        RouterModule
    ],
    declarations: []
})
export class AppRoutingModule { }