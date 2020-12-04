import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing/app-routing.module';


import { AppComponent } from './app.component';
import { ResumeComponent } from './resume/resume.component';
import { FrontPageComponent } from './front-page/front-page.component';



@NgModule({
  declarations: [
    AppComponent,
    ResumeComponent,
    FrontPageComponent,

  ],
  imports: [
    BrowserModule,
    FormsModule,
    BrowserAnimationsModule,
    AppRoutingModule,

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
